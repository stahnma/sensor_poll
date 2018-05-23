# Temp sensor poller for raspberry pi

This is an attempt to be a pure golang program to poll a temperature sensor on
a raspberry pi.  This is designed to work with a DHT22/AM2302 sensor that can
sense temperature and humidity.

This program uses go channels for timeouts, because the underlying libraries
that I'm using are buggy and will sometimes just hang or report mismatches. To
work around that (and not have to fall back to python), we simply time out if
the result isn't returned fast enough and bubble up other errors.

# Usage

build the program

    $ make deps
    $ make

From there, you can run `./sensor_poll`.

Currently, valid data is output on stdout. Errors and exceptions (of which
there appear to be several) are logged on stderr.


# Installation

If you want to install it, the program sensor_poll will be placed in
`/usr/local/bin`. A systemd unit will be placed in `/etc/systemd/system`. You can
then use standard systemd interactions for start/stop/journal query etc.

# GPIO PIN
This program assumes your GPIO sensor pin is pin 4. To change that, set the
GPIO_PIN environment variable prior to running.


# License
![WTFPL]( http://www.wtfpl.net/wp-content/uploads/2012/12/wtfpl-badge-4.png)

WTFPL
