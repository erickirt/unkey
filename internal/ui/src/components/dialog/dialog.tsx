"use client";

import * as DialogPrimitive from "@radix-ui/react-dialog";
import * as React from "react";

import { XMark } from "@unkey/icons";
import { cn } from "../../lib/utils";

const Dialog = DialogPrimitive.Root;

const DialogTrigger = DialogPrimitive.Trigger;

const DialogPortal = DialogPrimitive.Portal;

const DialogClose = DialogPrimitive.Close;

const DialogOverlay = React.forwardRef<
  React.ElementRef<typeof DialogPrimitive.Overlay>,
  React.ComponentPropsWithoutRef<typeof DialogPrimitive.Overlay> & {
    showCloseWarning?: boolean;
    onAttemptClose?: () => void;
  }
>(({ className, showCloseWarning = false, onAttemptClose, ...props }, ref) => (
  <DialogPrimitive.Overlay
    ref={ref}
    className={cn(
      "fixed inset-0 z-50 bg-black/30 backdrop-blur-sm data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0",
      className,
    )}
    {...props}
  />
));
DialogOverlay.displayName = DialogPrimitive.Overlay.displayName;

const DialogContent = React.forwardRef<
  React.ElementRef<typeof DialogPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof DialogPrimitive.Content> & {
    showCloseWarning?: boolean;
    onAttemptClose?: () => void;
    xButtonRef?: React.RefObject<HTMLButtonElement>;
  }
>(
  (
    { className, children, showCloseWarning = false, onAttemptClose, xButtonRef, ...props },
    ref,
  ) => {
    const handleCloseAttempt = React.useCallback(() => {
      // This handler is now only called when showCloseWarning is true
      if (showCloseWarning) {
        onAttemptClose?.();
      }
    }, [showCloseWarning, onAttemptClose]);

    // Common class names for both button types
    const buttonClassNames =
      "absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-accent text-muted-foreground z-[51]";

    return (
      <DialogPortal>
        <DialogOverlay showCloseWarning={showCloseWarning} onAttemptClose={handleCloseAttempt} />
        <DialogPrimitive.Content
          ref={ref}
          className={cn(
            "fixed left-[50%] top-[50%] z-50 grid w-full max-w-lg translate-x-[-50%] translate-y-[-50%] gap-4 border bg-background p-6 shadow-lg duration-200 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-left-1/2 data-[state=closed]:slide-out-to-top-[48%] data-[state=open]:slide-in-from-left-1/2 data-[state=open]:slide-in-from-top-[48%] sm:rounded-lg",
            className,
          )}
          onEscapeKeyDown={(e) => {
            if (showCloseWarning) {
              e.preventDefault();
              handleCloseAttempt();
            }
          }}
          onPointerDownOutside={(e) => {
            // Prevent closing only if warning is active and click is outside content
            if (showCloseWarning) {
              // Basic check: If the target is the overlay, it's handled there.
              // More robust checks might be needed depending on content complexity.
              const contentElement = (e.target as HTMLElement)?.closest('[role="dialog"]');
              if (!contentElement || contentElement !== e.currentTarget) {
                e.preventDefault();
                handleCloseAttempt();
              }
            }
          }}
          {...props}
        >
          {children}

          {/* Conditionally render the close button */}
          {showCloseWarning ? (
            <button
              ref={xButtonRef} // Attach ref only needed for anchoring the custom popover
              type="button"
              onClick={handleCloseAttempt} // Call attempt handler
              className={buttonClassNames}
              aria-label="Close dialog with confirmation"
            >
              <XMark size="md-regular" />
            </button>
          ) : (
            // Use DialogPrimitive.Close for standard behavior
            <DialogPrimitive.Close asChild>
              <button type="button" className={buttonClassNames} aria-label="Close dialog">
                <XMark size="md-regular" />
              </button>
            </DialogPrimitive.Close>
          )}
        </DialogPrimitive.Content>
      </DialogPortal>
    );
  },
);
DialogContent.displayName = DialogPrimitive.Content.displayName;

const DialogHeader = ({ className, ...props }: React.HTMLAttributes<HTMLDivElement>) => (
  <div className={cn("flex flex-col space-y-1.5 text-center sm:text-left", className)} {...props} />
);
DialogHeader.displayName = "DialogHeader";

const DialogFooter = ({ className, ...props }: React.HTMLAttributes<HTMLDivElement>) => (
  <div
    className={cn("flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2", className)}
    {...props}
  />
);
DialogFooter.displayName = "DialogFooter";

const DialogTitle = React.forwardRef<
  React.ElementRef<typeof DialogPrimitive.Title>,
  React.ComponentPropsWithoutRef<typeof DialogPrimitive.Title>
>(({ className, ...props }, ref) => (
  <DialogPrimitive.Title
    ref={ref}
    className={cn("text-lg font-semibold leading-none tracking-tight", className)}
    {...props}
  />
));
DialogTitle.displayName = DialogPrimitive.Title.displayName;

const DialogDescription = React.forwardRef<
  React.ElementRef<typeof DialogPrimitive.Description>,
  React.ComponentPropsWithoutRef<typeof DialogPrimitive.Description>
>(({ className, ...props }, ref) => (
  <DialogPrimitive.Description
    ref={ref}
    className={cn("text-sm text-content-subtle", className)}
    {...props}
  />
));
DialogDescription.displayName = DialogPrimitive.Description.displayName;

export {
  Dialog,
  DialogPortal,
  DialogOverlay,
  DialogClose,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogFooter,
  DialogTitle,
  DialogDescription,
};
