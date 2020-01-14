# Nyaa si watcher

A Go application that checks nyaa.si 's RSS feed for specified names. If a torrent's name appears in `watching` the id of the torrent (taken from nyaa.si/download/?.torrent)  isn't present in `resolved`, the torrent is downloadedd to a specified folder.

# Usage

1. Build

    - From withing the  directory run
    
        ```
        go build
        ```

2. TODO