# ASCII-ART-WEB

## Description
You can choose write a word, choose a banner file and a path, where you want to save the output.

If you want to display output to a terminal, here is usage:

    Usage: go run . [STRING] [BANNER]
    
    EX: go run . something standard

If you want to save your ouput to a file:

    Usage: go run . [STRING] [BANNER] [OPTION]

    EX: go run . something standard --output=<fileName.txt>

## How to run

    $ go run . "hello" standard | cat -e
     _              _   _          $
    | |            | | | |         $
    | |__     ___  | | | |   ___   $
    |  _ \   / _ \ | | | |  / _ \  $
    | | | | |  __/ | | | | | (_) | $
    |_| |_|  \___| |_| |_|  \___/  $
                                   $
                                   $
    $ go run . "Hello There!" shadow | cat -e
                                                                                          $
    _|    _|          _| _|                _|_|_|_|_| _|                                  $
    _|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   $
    _|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| $
    _|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|       $
    _|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| $
                                                                                          $
                                                                                          $
    $ go run . "Hello There!" thinkertoy | cat -e
                                                  $
    o  o     o o           o-O-o o                $
    |  |     | |             |   |                $
    O--O o-o | | o-o         |   O--o o-o o-o o-o $
    |  | |-' | | | |         |   |  | |-' |   |-' $
    o  o o-o o o o-o         o   o  o o-o o   o-o $
                                                  $
                                                  $
    $ go run . "hello" standard --output=banner.txt
    $ cat -e banner.txt
     _              _   _          $
    | |            | | | |         $
    | |__     ___  | | | |   ___   $
    |  _ \   / _ \ | | | |  / _ \  $
    | | | | |  __/ | | | | | (_) | $
    |_| |_|  \___| |_| |_|  \___/  $
                                   $
                                   $
    $
    $ go run . "Hello There!" shadow --output=banner.txt
    $ cat -e banner.txt
                                                                                             $
    _|    _|          _| _|                _|_|_|_|_| _|                                  _| $
    _|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
    _|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
    _|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
    _|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                             $
                                                                                             $
    $

