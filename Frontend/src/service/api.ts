import axios from "axios";
import { apiHost, apiPort } from "../utils/envConfig";

export const api = axios.create({
    baseURL: "http://" + apiHost + ":" + apiPort,
});

export default function authHeader() {
    const accessToken = sessionStorage.getItem('AuthAccessToken');
    if (accessToken !== null) {
      return { Authorization: 'Bearer ' + accessToken };
    } else {
      return {};
    }
}