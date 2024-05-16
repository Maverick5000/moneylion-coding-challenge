import { IContentCard } from "@/types/IContentCard";
import HttpClient from "@/utils/httpClient";
import { AxiosError, AxiosResponse } from "axios";

/**
 * @returns {IcontentCard[]}
 * @description Hook to fetch feed data from a specified API endpoint.
 */
const httpClient = new HttpClient("http://api:8080", {
  "Content-Type": "application/json",
});

const useGetFeedData = async () => {
  const response = await httpClient
    .get("/feed")
    .then((response: AxiosResponse) => {
      return response.data as IContentCard[];
    })
    .catch((error: AxiosError) => {
      throw error;
    });
  return response;
};

export default useGetFeedData;
