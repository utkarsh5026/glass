import axios, { AxiosInstance, AxiosRequestConfig } from "axios";
import store from "../store/store";

const baseURL = "http://localhost:8080/api";

/**
 * Creates an Axios instance with predefined configuration.
 * @constant
 * @type {AxiosInstance}
 */
const axiosInstance: AxiosInstance = axios.create({
  baseURL,
  timeout: 10000, // 10 seconds
  headers: {
    "Content-Type": "application/json",
  },
});

/**
 * Intercepts requests to add authorization token if available.
 */
axiosInstance.interceptors.request.use(
  (config) => {
    const state = store.getState();
    const token = state.auth.token;
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }
    return config;
  },
  (error) =>
    Promise.reject(error instanceof Error ? error : new Error(String(error)))
);

/**
 * Makes an API call using the configured Axios instance.
 * @template T - The expected type of the response data.
 * @param {AxiosRequestConfig} config - The configuration for the API request.
 * @returns {Promise<T>} A promise that resolves with the response data.
 * @throws {Error} Throws an error if the API call fails.
 */
export const apiCall = async <T>(config: AxiosRequestConfig): Promise<T> => {
  try {
    const response = await axiosInstance(config);
    return response.data;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      throw error.response?.data || error.message;
    }
    throw error;
  }
};

export default axiosInstance;
