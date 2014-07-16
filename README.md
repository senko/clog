# clog - Colorful logger for Go with support for different log levels

[![Build Status](https://travis-ci.org/senko/clog.svg?branch=master)](https://travis-ci.org/senko/clog?branch=master)

Five log levels are predefined: DEBUG, INFO, WARNING, ERROR and PANIC. The
logger will only output log messages with level equal to or higher than
limit specified in setup. The messages can optionally be shown in color
(turned off by default).

All messages are shown with a RFC3339 timestamp.

The logger can be setup directly using Setup(). Alternatively, using
SetupFromEnv(), the settings can be picked from environment variables
LOG_LEVEL (should be one of the predefined levle names) and
LOG_COLOR (should be "true" or "false").

The logger provides Log() function which takes a level, and a message. The
convenience functions Debug(), Info(), Warning(), Error() and Panic() are
also provided. There are also variants of these functions which support
passing a format string and arguments instead of a single message string:
Logf(), Debugf(), Infof(), Warningf(), Errorf() and Panicf().

When logging a message with a PANIC level, the logger will raise a panic
with the specified message immediately after logging it.

The output by default goes to os.Stderr. This can be changed by using
SetOutput(). Note that SetOutput() must be called after Setup() (or
SetupFromEnv()).

Example use:

    import "clog"

    clog.Setup(clog.WARNING, true)
    clog.Debug("hello world!")
    clog.Warning("Hello")
    clog.Panicf("The end is %s!", "nigh")


## License

Copyright (C) 2014. Senko Rašić.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
