"use client";
import useTruncatedElement from "@/hooks/useTruncatedElement";
import type { IContentCard } from "@/types/IContentCard";
import Image from "next/image";
import { memo, useRef, useState } from "react";
import Comments from "./Comments";

const ContentCard = ({ data }: { data: IContentCard }) => {
  const [showComments, setShowComments] = useState(false);
  const ref = useRef(null);
  const { isTruncated, isReadingMore, setIsReadingMore } = useTruncatedElement({
    ref,
  });

  return (
    <div className="card card-compact w-96 xl:w-[28rem] bg-base-100 shadow-xl prose">
      <Image
        src={data.imageUri}
        placeholder="data:image/png;base64,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"
        alt="Feed"
        width={0}
        height={0}
        sizes="100vw"
        className="object-contain bg-base-content w-full h-64"
        style={{
          borderRadius: "0.5rem 0.5rem 0 0",
          marginBottom: 0,
        }}
      />

      <div className="card-body gap-0">
        <div className="flex justify-between">
          <h2 className="card-title mt-0">{data.title}</h2>
          <h2 className="card-title mt-0">{data.author}</h2>
        </div>
        <h4 className="heading-4 mt-0">{data.subtitle}</h4>
        <p
          ref={ref}
          className={`m-0 break-words ${!isReadingMore && "line-clamp-3"}`}
        >
          {data.body}
        </p>
        {isTruncated && !isReadingMore && (
          <button
            className="btn btn-link btn-xs ml-auto"
            onClick={() => setIsReadingMore(true)}
          >
            Read more
          </button>
        )}
        <div className="card-actions justify-start">
          {!showComments && (
            <button
              className="btn btn-link btn-xs"
              onClick={() => setShowComments(true)}
            >
              Comments
            </button>
          )}
        </div>
        {showComments && <Comments comments={data.comments} />}
      </div>
    </div>
  );
};

export default memo(ContentCard);
