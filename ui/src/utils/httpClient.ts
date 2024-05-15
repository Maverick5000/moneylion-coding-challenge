import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
/**
 * A class to handle HTTP requests using Axios.
 * This class provides methods for making GET, POST, PUT, and DELETE requests.
 *
 * Usage:
 *
 * const httpClient = new HttpClient('https://api.example.com', { 'Authorization': 'Bearer token' });
 *
 * // Making a GET request
 * const response = await httpClient.get('/endpoint');
 *
 * // Making a POST request
 * const response = await httpClient.post('/endpoint', { data: 'sample' });
 *
 * // Making a PUT request
 * const response = await httpClient.put('/endpoint', { data: 'sample' });
 *
 * // Making a DELETE request
 * const response = await httpClient.delete('/endpoint');
 *
 * Methods:
 * - get<T>(url: string, config?: AxiosRequestConfig): Promise<AxiosResponse<T>>
 * - post<T>(url: string, data?: any, config?: AxiosRequestConfig): Promise<AxiosResponse<T>>
 * - put<T>(url: string, data?: any, config?: AxiosRequestConfig): Promise<AxiosResponse<T>>
 * - delete<T>(url: string, config?: AxiosRequestConfig): Promise<AxiosResponse<T>>
 */
class HttpClient {
  private readonly client: AxiosInstance;

  constructor(baseURL?: string, headers?: Record<string, string>) {
    this.client = axios.create({
      baseURL: baseURL || process.env.BASE_API_URL,
      headers,
    });
  }

  public async get<T>(
    url: string,
    config?: AxiosRequestConfig,
  ): Promise<AxiosResponse<T>> {
    const response: AxiosResponse<T> = await this.client.get(url, config);
    return response;
  }

  public async post<T>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig,
  ): Promise<AxiosResponse<T>> {
    const response: AxiosResponse<T> = await this.client.post(
      url,
      data,
      config,
    );
    return response;
  }

  public async put<T>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig,
  ): Promise<AxiosResponse<T>> {
    const response: AxiosResponse<T> = await this.client.put(url, data, config);
    return response;
  }

  public async delete<T>(
    url: string,
    config?: AxiosRequestConfig,
  ): Promise<AxiosResponse<T>> {
    const response: AxiosResponse<T> = await this.client.delete(url, config);
    return response;
  }
}

export default HttpClient;
