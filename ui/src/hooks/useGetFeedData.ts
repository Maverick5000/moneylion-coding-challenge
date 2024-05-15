import { IContentCard } from "@/types/IContentCard";
import HttpClient from "@/utils/httpClient";
import { AxiosError, AxiosResponse } from "axios";

/**
 * @returns {IcontentCard[]}
 * @description Hook to fetch feed data from a specified API endpoint.
 */
const httpClient = new HttpClient("https://stoplight.io", {
  "Content-Type": "application/json",
});

const useGetFeedData = async () => {
  const response = await httpClient
    .get("/mocks/engine/fullstack-spec/52502230/content")
    .then((response: AxiosResponse) => {
      return response.data.contentCards as IContentCard[];
    })
    .catch((error: AxiosError) => {
      throw error;
    });
  return response;
};

export default useGetFeedData;
