import ContentCard from "@/components/ContentCard";
import useGetFeedData from "@/hooks/useGetFeedData";
import type { IContentCard } from "@/types/IContentCard";

export default async function Home() {
  const data = await useGetFeedData();

  return (
    <div className="flex flex-col items-center overflow-scroll w-screen h-[calc(100%-4rem)] md:h-[calc(100%-3rem)] gap-4 py-2 bg-base-300">
      {data.map((card: IContentCard) => (
        <ContentCard key={card.id + Math.random()} data={card} />
      ))}
    </div>
  );
}
