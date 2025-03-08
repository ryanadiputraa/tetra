"use client";

import { useState, useCallback, useEffect } from "react";

export const useMediaQuery = (query: string) => {
  const [matches, setMatches] = useState(false);

  const updateMatch = useCallback(() => {
    setMatches(window.matchMedia(query).matches);
  }, [query]);

  useEffect(() => {
    updateMatch(); // Run on mount
    window.addEventListener("resize", updateMatch);
    return () => window.removeEventListener("resize", updateMatch);
  }, [updateMatch]);

  return matches;
};
