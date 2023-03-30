# Cleanium

## When Private/Incognito just won't do!

This is a wrapper for Chrome/Chromium which takes advantage of the feature to specify a specific data directory.  Each time this app is launched it:
* creates a new temporary directory
* opens chrome or chromium with the temporary directory as it's data storage directory
* skips the "first run" BS just as a matter of convenience

The end result is each time you click/open Cleanium you get a brand new fresh browser with zero history, cache, cookies, extensions, etc.

### Settings

A development machine generally has lots of browsers installed and may have multiple versions of Chrome and Chromium installed.  Currently the only setting is the `BROWSER_USER_DIR` environment variable.  Specifying a path here to a version of chrome/chromium on your hard drive enables the use of multiple/different browser versions as needed.  

### Roadmap

* Settings File: In addition to the environment variable method of specifying a browser, adding support for a setting file in the users home directory would enable specifying which browser to execute globally.
* Cross OS Support: This was written to fill an immidiate need for developers and QA _using macs_.  The code is structured with the forethought of eventually supporting windows and linux as well.  (ie help wanted)
* Notifications: currently is does a horrible job of logging or communicating errors.  This is not the biggest issue as all this application does is prepare and hand over control to the version Chrome or Chromium already installed.  However, it would be better to bundled in (OS-specific) notifications to help with the initial setup.