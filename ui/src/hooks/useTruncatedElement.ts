import { RefObject, useLayoutEffect, useState } from "react";

interface UseTruncatedElementProps {
  ref: RefObject<HTMLElement>;
}

interface UseTruncatedElementReturn {
  isTruncated: boolean;
  isReadingMore: boolean;
  setIsReadingMore: (readingMore: boolean) => void;
}

/**
 * @param {UseTruncatedElementProps} props
 * @returns {UseTruncatedElementReturn}
 * @description Hook to detect if a text element is truncated
 */
const useTruncatedElement = ({
  ref,
}: UseTruncatedElementProps): UseTruncatedElementReturn => {
  const [isTruncated, setIsTruncated] = useState(false);
  const [isReadingMore, setIsReadingMore] = useState(false);

  useLayoutEffect(() => {
    const { offsetHeight, scrollHeight } = ref.current || {};

    if (offsetHeight && scrollHeight && offsetHeight < scrollHeight) {
      setIsTruncated(true);
    } else {
      setIsTruncated(false);
    }
  }, [ref]);

  return {
    isTruncated,
    isReadingMore,
    setIsReadingMore,
  };
};

export default useTruncatedElement;
