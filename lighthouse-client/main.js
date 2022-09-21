const puppeteer = require('puppeteer');
const lighthouse = require('lighthouse');
const {URL} = require('url');

/**
 * analyze the given url with chrome's lighthouse
 *
 * @param url   url of website to analyze with lighthouse
 * @param chromeExec path to chrome executable (in docker) or empty string
 */
async function runLighthouse(url, chromeExec) {

    if (chromeExec === "") {
        chromeExec = null;
    }

    const browser = await puppeteer.launch({
        headless: true,
        defaultViewport: null,
        executablePath: chromeExec,
    });

    const {lhr} = await lighthouse(url, {
        port: (new URL(browser.wsEndpoint())).port,
        output: 'json',
        logLevel: 'error',
    });

    process.stdout.write(JSON.stringify(lhr));

    await browser.close();
}


const url = process.argv[2];
const chromeExec = process.argv[3];
runLighthouse(url, chromeExec).then();
