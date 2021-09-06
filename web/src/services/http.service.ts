import axios from "axios";
import type {AxiosAdapter} from  'axios'

export function setHttpHeader(key: string, value: string): AxiosAdapter {
    axios.defaults.headers[key] = value;
    return axios;
}

export function getClient() {
    return axios;
}