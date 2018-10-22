# Nyaa si watcher

A Go application that checks nyaa.si 's RSS feed for specified names. If a torrent's name appears in `watching` the id of the torrent (taken from nyaa.si/download/?.torrent)  isn't present in `resolved`, the file is added using `deluge-console add`.

# Requirements

Currently the program only works with Deluge, as it uses `deluge-console` to add new files. Alternatively it is possible to change the `addTorrent` function in `torrentoptions/contenthandler.go` to make it work with something else.

# Usage

1. Build

    - From withing the  directory run
    
        ```
        go build
        ```

2. Prepare the `watching` and `resolved` files. By default the program will look in `/var/lib/nyaa-si-watcher/`, but this can be changed by passing `-confDir`. Each line will contain one name in `watching` or one id in `resolved`.
Prepare `announcemails` - an email per line to send emails to using `mailutils`, notifying about changes.

3. Run it.
    ```
    nyaa-si-watcher
    ```
    
    The program checks every 30 seconds for updates, and shows the watched titles are shown in a webpage at port 80. (This will be used for settings in the future.)
