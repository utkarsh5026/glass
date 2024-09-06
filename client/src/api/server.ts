import axios, { AxiosInstance, AxiosRequestConfig } from "axios";

const baseURL = "http://localhost:8080/api";

const axiosInstance: AxiosInstance = axios.create({
  baseURL,
  timeout: 10000, // 10 seconds
  headers: {
    "Content-Type": "application/json",
  },
});

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
