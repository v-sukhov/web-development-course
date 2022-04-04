import Response from "./Response";

const BASE_URL = 'https://localhost:8443/api'

const LOGIN_URL = BASE_URL + '/login'

const storagename = 'resAggData'

function Datahub(){

    this.authToken = "";
    this.login = "";

    this.requestService = (serviceUrl, method, body, callback) => {

        let headers = {
            "Content-Type": "application/json"
        }

        if (serviceUrl !== LOGIN_URL){
            headers["Authorization"] = this.authToken
        }
        
        const request = {
            method,
            headers
        }

        if (method === "POST"){
            request.body = JSON.stringify(body)
        }

        fetch(BASE_URL + serviceUrl, request).then(response => {
            response.json().then(json => {
                if (json.success) {
                    if (serviceUrl === LOGIN_URL && response.ok) {
                        json.data = {...json.data, token: json.token}
                    }
                    callback(new Response(true, json.data))
                }
                else {
                    callback(new Response(false, json.message))
                }
            })
        })
    }

    this.doLogin = (login, password, callback) => {
        const body = {
            login,
            password
        }

        this.requestService(LOGIN_URL, "POST", body, 
            (response) => {
                if(response.success){
                    this.login = login;
                    this.authToken = response.data.token;
                    localStorage.setItem(storagename,JSON.stringify({authToken: this.authToken, login: login}));
                }
                callback(response);
            }
        );
    }

    this.clearAuthentication = () => {
        this.authToken = "";
        this.login = "";
        localStorage.removeItem(storagename);
    }

    this.getAuthenticaion = () => {
        let storageItem = localStorage.getItem(storagename);
        if(storageItem === null){
            return false;
        }
        else{
            let storageObj = JSON.parse(storageItem);
            if(storageObj.login && storageObj.authToken){
                this.login = storageObj.login;
                this.authToken = storageObj.authToken;
                return true;
            }
            else{
                return false;
            }
        }
    }
}



const datahub = new Datahub();

export default datahub;
