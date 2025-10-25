import React, { useRef, useState, useEffect } from 'react';
import ArrowRight from '../icons/ArrowRight';
import ArrowLeft from '../icons/ArrowLeft';

export type Service = {
  id: string;
  name: string;
  /** Icon as a ReactNode (SVG JSX, <img />, or any element) */
  icon: React.ReactNode;
};

type Props = {
  services: Service[];
  /** how many items are visible on large screens (tailwind breakpoints handled via CSS); used for scrolling distance */
  visibleCount?: number;
};

export default function ServicesCarousel({ services, visibleCount = 4 }: Props) {
  const ref = useRef<HTMLDivElement | null>(null);
  const [canScrollLeft, setCanScrollLeft] = useState(false);
  const [canScrollRight, setCanScrollRight] = useState(false);

  useEffect(() => {
    const el = ref.current;
    if (!el) return;
    const update = () => {
      setCanScrollLeft(el.scrollLeft > 0);
      setCanScrollRight(el.scrollLeft + el.clientWidth < el.scrollWidth - 1);
    };
    update();
    el.addEventListener("scroll", update);
    window.addEventListener("resize", update);
    return () => {
      el.removeEventListener("scroll", update);
      window.removeEventListener("resize", update);
    };
  }, [services]);

  const scrollByWidth = (dir: "left" | "right") => {
    const el = ref.current;
    if (!el) return;
    // scroll by the visible width (so it pages by columns)
    const distance = Math.max(el.clientWidth, (el.clientWidth / visibleCount) * Math.floor(visibleCount));
    el.scrollBy({ left: dir === "right" ? distance : -distance, behavior: "smooth" });
  };

  return (
    <div className="relative">
      {/* Prev button */}
      <button
        type="button"
        className={`absolute left-0 top-1/2 -translate-y-1/2 z-10 p-2 rounded-full shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 backdrop-blur-sm hover:scale-105 transition-transform disabled:opacity-40 disabled:cursor-not-allowed`}
        aria-label="Scroll para a esquerda"
        onClick={() => scrollByWidth("left")}
        disabled={!canScrollLeft}
      >
        <ArrowRight className="w-4 h-4 text-gray-700"/>
      </button>

      {/* Next button */}
      <button
        type="button"
        className={`absolute right-0 top-1/2 -translate-y-1/2 z-10 p-2  focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 backdrop-blur-sm hover:scale-105 transition-transform disabled:opacity-40 disabled:cursor-not-allowed`}
        aria-label="Scroll para a direita"
        onClick={() => scrollByWidth("right")}
        disabled={!canScrollRight}
      >
        <ArrowLeft className="w-4 h-4 text-gray-700"/>
      </button>

      <div
        ref={ref}
        className="overflow-x-auto scrollbar-none scroll-smooth py-4"
        role="list"
        aria-label="Serviços disponíveis"
        tabIndex={0}
      >
        <div
          className="inline-flex gap-4 px-6"
          style={{
            // make items shrink-to-fit and allow snapping
            whiteSpace: "nowrap",
          }}
        >
          {services.map((s) => (
            <div
              key={s.id}
              role="listitem"
              className="flex-none w-28 sm:w-36 md:w-40 lg:w-44 xl:w-48 text-center"
            >
              <button
                className={`w-full h-full flex flex-col items-center gap-2 p-2 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 hover:shadow-lg transition-shadow`}
                aria-label={s.name}
                type="button"
              >
                <div className="w-14 h-14 sm:w-16 sm:h-16 md:w-18 md:h-18 rounded-full flex items-center justify-center bg-white shadow-sm ring-1 ring-gray-100">
                  {/* Icon container: render whatever the user provided */}
                  <span className="inline-flex items-center justify-center" aria-hidden>
                    {s.icon}
                  </span>
                </div>
                <span className="mt-1 text-xs sm:text-sm md:text-base text-gray-800 font-medium truncate">{s.name}</span>
              </button>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}


