package tests

import (
	"context"
	"github.com/chromedp/chromedp"
	"testing"
	"time"
)

// TestSearchEndToEnd tests the search functionality end to end. It will be skipped by unit tests.
// Some steps are performed:
// - open the browser
// - navigate to the search page
// - enter a search term
// - click the search button
// - check if result is displayed
// - click the read more button
// - check if the chapter content is displayed
func TestSearchEndToEnd(t *testing.T) {

	// this will be skipped by unit tests
	if testing.Short() {
		t.Skip("skipping end to tend test")
	}

	// headless is disabled for local development
	opts := append(chromedp.DefaultExecAllocatorOptions[3:],
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.IgnoreCertErrors,
		//chromedp.Headless,
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create chrome instance
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// url is hardcoded for local development, it could come from environment variable as well
	if err := chromedp.Run(
		ctx,
		chromedp.Navigate("http://localhost:1234"),
		chromedp.WaitVisible("//button[@type='submit'][text()='Search']"),
	); err != nil {
		t.Error(err.Error())
	}

	if err := chromedp.Run(
		ctx,
		chromedp.SendKeys("//input[@id='query']", "London SIXTH CLAReNCE WINCHESTER"),
	); err != nil {
		t.Error(err.Error())
	}

	if err := chromedp.Run(
		ctx,
		chromedp.Click("//button[@type='submit'][text()='Search']"),
	); err != nil {
		t.Error(err.Error())
	}

	// todo: this is for debugging, so we can see the automated inputs
	time.Sleep(3 * time.Second)

	if err := chromedp.Run(
		ctx,
		chromedp.WaitVisible("//div[@class='query-result']"),
	); err != nil {
		t.Error(err.Error())
	}

	if err := chromedp.Run(
		ctx,
		chromedp.Click("(//a[text()='read more'])[1]"),
	); err != nil {
		t.Error(err.Error())
	}
	if err := chromedp.Run(
		ctx,
		chromedp.WaitVisible("//div[@class='chapter-content']"),
	); err != nil {
		t.Error(err.Error())
	}
	// todo: this is for debugging, so we can see the automated inputs
	time.Sleep(3 * time.Second)

}
