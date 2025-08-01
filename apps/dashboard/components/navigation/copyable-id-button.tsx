"use client";
import { Button, CopyButton } from "@unkey/ui";
import { useRef } from "react";

type CopyableIDButtonProps = {
  value: string;
  className?: string;
};

export const CopyableIDButton = ({ value, className = "" }: CopyableIDButtonProps) => {
  const textRef = useRef<HTMLDivElement>(null);
  const pressTimer = useRef<NodeJS.Timeout | null>(null);
  const copyButtonRef = useRef<HTMLButtonElement>(null);

  const handleMouseDown = () => {
    // Start a long-press timer
    pressTimer.current = setTimeout(() => {
      // For long-press, select the text
      if (textRef.current) {
        const range = document.createRange();
        range.selectNodeContents(textRef.current);
        const selection = window.getSelection();
        if (selection) {
          selection.removeAllRanges();
          selection.addRange(range);
        }
      }
    }, 500);
  };

  const handleMouseUp = () => {
    // Clear the timer if mouse is released before long-press threshold
    if (pressTimer.current) {
      clearTimeout(pressTimer.current);
      pressTimer.current = null;
    }
  };

  const handleMouseLeave = () => {
    // Clear the timer if mouse leaves the button
    if (pressTimer.current) {
      clearTimeout(pressTimer.current);
      pressTimer.current = null;
    }
  };

  const handleClick = (e: React.MouseEvent) => {
    // Only handle click if it wasn't a long press
    if (window.getSelection()?.toString()) {
      // If text is selected, don't trigger the copy
      e.stopPropagation();
    } else {
      // Programmatically click the CopyButton if text isn't selected
      copyButtonRef.current?.click();
    }
  };

  return (
    <Button
      variant="outline"
      size="md"
      className={`text-xs font-mono font-medium ph-no-capture h-7 bg-grayA-2 hover:bg-grayA-3 w-[190px] ${className}`}
      onMouseDown={handleMouseDown}
      onMouseUp={handleMouseUp}
      onMouseLeave={handleMouseLeave}
      onClick={handleClick}
    >
      <div className="flex gap-2 items-center justify-between w-full min-w-0">
        <div ref={textRef} className="select-text truncate min-w-0 flex-1">
          {value}
        </div>
        <CopyButton
          variant="ghost"
          value={value}
          ref={copyButtonRef}
          toastMessage={value}
          className="pointer-events-none flex-shrink-0"
        />
      </div>
    </Button>
  );
};
