package main

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type pageHandler func(ctx context.Context, frameID cdp.FrameID) error

func main() {
	exCtx, cancel := newExecContext()
	defer cancel()

	newPageExec(exCtx, func(tabCtx context.Context, frameID cdp.FrameID) error {
		page.
			SetDocumentContent(frameID, `
				<html>
					<head></head>
					<body>
						<h1>Test</h1>
					</body>
				</html>
			`).
			Do(tabCtx)

		return nil
	})
}

func newExecContext() (context.Context, context.CancelFunc) {
	allocatorOptions := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
	)

	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		allocatorOptions...,
	)

	return ctx, cancel
}

func newPageExec(execCtx context.Context, handler pageHandler) {
	newTabCtx, cancel := chromedp.NewContext(execCtx)
	chromedp.Run(newTabCtx, taskHandler(handler))
	cancel()
}

func taskHandler(handler pageHandler) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)

			if err != nil {
				return err
			}

			err = handler(ctx, frameTree.Frame.ID)

			if err != nil {
				return err
			}

			return nil
		}),
	}
}
