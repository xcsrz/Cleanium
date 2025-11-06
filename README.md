# Cleanium

![Cleanium Icon](icon/Icon.iconset/icon_256x256.png)

## When Private/Incognito just won't do!

This is a wrapper for Chrome/Chromium which takes advantage of the feature to specify a specific data directory.  Each time this app is launched it:
* creates a new temporary directory
* opens chrome or chromium with the temporary directory as it's data storage directory
* skips the "first run" BS just as a matter of convenience

The end result is each time you click/open Cleanium you get a brand new fresh browser with zero history, cache, cookies, extensions, etc.

### Installation

1. Download the lastest version from the [Releases page](https://github.com/xcsrz/Cleanium/releases)
    - The `arm64` file is correct for all newer macs
    - The `amd64` file is the one if you're on an older "intel" mac
2. Unzip the file and move the `Cleanium` app (or `Cleanium.app` on the command line) to your application folder of choice
    - `/Applications` is the defacto location
    - `~/Applications` is another good choice especially if you have multiple users and user accounts on the machine
3. Run the following command to let macos know you trust the app and want to make it runnable by executing this in your command line.

```
xattr -d com.apple.quarantine /path/to/Cleanium.app
```

### Custom Builds

Cleanium can be customized to create specialized versions for specific use cases. The build system supports the following variables:

#### APP_NAME
Specifies a custom name for the application bundle. This changes both the `.app` bundle name and the executable name.

**Example:**
```bash
APP_NAME=MyCustomBrowser make build
```

#### APP_URL
Sets a default URL that will be opened when the browser launches. This is useful for creating purpose-built browsers for specific web applications.

**Example:**
```bash
APP_NAME=MailBrowser APP_URL=https://mail.google.com make build
```

#### APP_ICON
Specifies a custom `.icns` icon file to use instead of the default Cleanium icon. The icon will be copied as `Icon.icns` in the app bundle.

**Example:**
```bash
APP_NAME=WorkBrowser APP_ICON=/path/to/custom-icon.icns make build
```

#### Combined Example
Create a fully customized browser:
```bash
APP_NAME=AnonyGPT APP_URL=https://chatgpt.com APP_ICON=~/Downloads/AIbot.icns make build-release
```

This creates a specialized cleanium browser with the same zero history or connection to your normal Chome browser as Cleanium, but named "AnonyGPT" and opens ChatGPT on each launch in a new window with a custom AI themed icon to help you tell the difference.

### Settings

A development machine generally has lots of browsers installed and may have multiple versions of Chrome and Chromium installed.  Currently the only setting is the `BROWSER_USER_DIR` environment variable.  Specifying a path here to a version of chrome/chromium on your hard drive enables the use of multiple/different browser versions as needed.  

### Roadmap

* Settings File: In addition to the environment variable method of specifying a browser, adding support for a setting file in the users home directory would enable specifying which browser to execute globally.
* Cross OS Support: This was written to fill an immidiate need for developers and QA _using macs_.  The code is structured with the forethought of eventually supporting windows and linux as well.  (ie help wanted)
* Notifications: currently is does a horrible job of logging or communicating errors.  This is not the biggest issue as all this application does is prepare and hand over control to the version Chrome or Chromium already installed.  However, it would be better to bundled in (OS-specific) notifications to help with the initial setup.