import axios, { AxiosError, AxiosResponse } from "axios";

import { API_URL, SERVER_ERR_MSG } from "@/constant";

export const fetcher = axios.create({
  baseURL: API_URL,
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
  },
});

fetcher.interceptors.response.use(
  (response: AxiosResponse) => response,
  (error: AxiosError) => {
    if (error.response) {
      if (error.response.data) {
        return Promise.reject(error.response.data);
      }
    }
    return Promise.reject(new Error(SERVER_ERR_MSG));
  },
);
