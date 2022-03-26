# Contents

- [Audio and Music](#audio-and-music)
- [Authentication and OAuth](#authentication-and-oauth)
- [Blockchain](#blockchain)
- [Bot Building](#bot-building)
- [Build Automation](#build-automation)
- [Command Line](#command-line)
- [Advanced Console UIs](#advanced-console-uis)
- [Standard CLI](#standard-cli)
- [Configuration](#configuration)
- [Continuous Integration](#continuous-integration)
- [CSS Preprocessors](#css-preprocessors)
- [Data Structures](#data-structures)
- [Database](#database)
- [Date and Time](#date-and-time)
- [Distributed Systems](#distributed-systems)
- [Dynamic DNS](#dynamic-dns)
- [Email](#email)
- [Embeddable Scripting Languages](#embeddable-scripting-languages)
- [Error Handling](#error-handling)
- [File Handling](#file-handling)
- [Financial](#financial)
- [Forms](#forms)
- [Functional](#functional)
- [Game Development](#game-development)
- [Generators](#generators)
- [Geographic](#geographic)
- [Go Compilers](#go-compilers)
- [Goroutines](#goroutines)
- [GUI](#gui)
- [Hardware](#hardware)
- [Images](#images)
- [IoT (Internet of Things)](#iot-internet-of-things)
- [Job Scheduler](#job-scheduler)
- [JSON](#json)
- [Logging](#logging)
- [Machine Learning](#machine-learning)
- [Messaging](#messaging)
- [Microsoft Office](#microsoft-office)
- [Microsoft Excel](#microsoft-excel)
- [Miscellaneous](#miscellaneous)
- [Dependency Injection](#dependency-injection)
- [Project Layout](#project-layout)
- [Strings](#strings)
- [Uncategorized](#uncategorized)
- [Natural Language Processing](#natural-language-processing)
- [Language Detection](#language-detection)
- [Morphological Analyzers](#morphological-analyzers)
- [Slugifiers](#slugifiers)
- [Tokenizers](#tokenizers)
- [Translation](#translation)
- [Transliteration](#transliteration)
- [Networking](#networking)
- [HTTP Clients](#http-clients)
- [OpenGL](#opengl)
- [ORM](#orm)
- [Package Management](#package-management)
- [Performance](#performance)
- [Query Language](#query-language)
- [Resource Embedding](#resource-embedding)
- [Science and Data Analysis](#science-and-data-analysis)
- [Security](#security)
- [Serialization](#serialization)
- [Server Applications](#server-applications)
- [Stream Processing](#stream-processing)
- [Template Engines](#template-engines)
- [Testing](#testing)
- [Text Processing](#text-processing)
- [Formatters](#formatters)
- [Markup Languages](#markup-languages)
- [Regular Expressions](#regular-expressions)
- [Sanitation](#sanitation)
- [Scrapers](#scrapers)
- [RSS](#rss)
- [Third-party APIs](#third-party-apis)
- [Utilities](#utilities)
- [UUID](#uuid)
- [Validation](#validation)
- [Version Control](#version-control)
- [Video](#video)
- [Web Frameworks](#web-frameworks)
- [Actual middlewares](#actual-middlewares)
- [Libraries for creating HTTP middlewares](#libraries-for-creating-http-middlewares)
- [Routers](#routers)
- [WebAssembly](#webassembly)
- [Windows](#windows)
- [XML](#xml)
- [Zero Trust](#zero-trust)
- [Code Analysis](#code-analysis)
- [Editor Plugins](#editor-plugins)
- [Go Generate Tools](#go-generate-tools)
- [Go Tools](#go-tools)
- [Software Packages](#software-packages)
- [DevOps Tools](#devops-tools)
- [Other Software](#other-software)
# Audio and Music

| Repository | Description | Stars |
|:---|:---|:---:|
| [mewkiz/flac](https://github.com/mewkiz/flac) | Package flac provides access to FLAC (Free Lossless Audio Codec) streams. | 201 | 
| [Comcast/gaad](https://github.com/Comcast/gaad) | GAAD (Go Advanced Audio Decoder) | 97 | 
| [DylanMeeus/GoAudio](https://github.com/DylanMeeus/GoAudio) | Go tools for audio processing & creation 🎶 | 216 | 
| [dh1tw/gosamplerate](https://github.com/dh1tw/gosamplerate) | Go Bindings for libsamplerate | 17 | 
| [bogem/id3v2](https://github.com/bogem/id3v2) | 🎵 ID3 decoding and encoding library for Go | 237 | 
| [gen2brain/malgo](https://github.com/gen2brain/malgo) | Mini audio library | 169 | 
| [tosone/minimp3](https://github.com/tosone/minimp3) | Decode mp3 base on https://github.com/lieff/minimp3 | 75 | 
| [go-mix/mix](https://github.com/go-mix/mix) | Sequence-based Go-native audio mixer for music apps | 153 | 
| [go-music-theory/music-theory](https://github.com/go-music-theory/music-theory) | Go models of Note, Scale, Chord and Key | 387 | 
| [hajimehoshi/oto](https://github.com/hajimehoshi/oto) | ♪ A low-level library to play sound on multiple platforms ♪ | 968 | 
| [gordonklaus/portaudio](https://github.com/gordonklaus/portaudio) | Go bindings for the PortAudio audio I/O library | 495 | 
| [rakyll/portmidi](https://github.com/rakyll/portmidi) | Go bindings for libportmidi | 265 | 


# Authentication and OAuth

| Repository | Description | Stars |
|:---|:---|:---:|
| [volatiletech/authboss](https://github.com/volatiletech/authboss) | The boss of http auth. | 3014 | 
| [essentialkaos/branca](https://github.com/essentialkaos/branca) | Authenticated encrypted API tokens (IETF XChaCha20-Poly1305 AEAD) for Golang | 31 | 
| [casbin/casbin](https://github.com/casbin/casbin) | An authorization library that supports access control models like ACL, RBAC, ABAC in Golang | 11563 | 
| [mengzhuo/cookiestxt](https://github.com/mengzhuo/cookiestxt) | cookiestxt implement parser of cookies txt format | 12 | 
| [dimuska139/go-email-normalizer](https://github.com/dimuska139/go-email-normalizer) | Golang library for providing a canonical representation of email address. | 40 | 
| [shaj13/go-guardian](https://github.com/shaj13/go-guardian) | Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication. | 354 | 
| [square/go-jose](https://github.com/square/go-jose) | An implementation of JOSE standards (JWE, JWS, JWT) in Go | 1899 | 
| [dghubble/gologin](https://github.com/dghubble/gologin) | Go login handlers for authentication providers (OAuth1, OAuth2) | 1503 | 
| [mikespook/gorbac](https://github.com/mikespook/gorbac) | goRBAC provides a lightweight role-based access control (RBAC) implementation in Golang. | 1256 | 
| [markbates/goth](https://github.com/markbates/goth) | Package goth provides a simple, clean, and idiomatic way to write authentication packages for Go web applications. | 3566 | 
| [abraithwaite/jeff](https://github.com/abraithwaite/jeff) | 🍍Jeff provides the simplest way to manage web sessions in Go. | 237 | 
| [pascaldekloe/jwt](https://github.com/pascaldekloe/jwt) | JSON Web Token library | 282 | 
| [cristalhq/jwt](https://github.com/cristalhq/jwt) | Safe, simple and fast JSON Web Tokens for Go | 323 | 
| [adam-hanna/jwt-auth](https://github.com/adam-hanna/jwt-auth) | This package provides json web token (jwt) middleware for goLang http servers | 213 | 
| [tarent/loginsrv](https://github.com/tarent/loginsrv) | JWT login microservice with plugable backends such as OAuth2, Google, Github, htpasswd, osiam, .. | 1849 | 
| [golang/oauth2](https://github.com/golang/oauth2) | Go OAuth2 | 4087 | 
| [openshift/osin](https://github.com/openshift/osin) | Golang OAuth2 server library | 1720 | 
| [RijulGulati/otpgen](https://github.com/RijulGulati/otpgen) | Library to generate TOTP/HOTP codes | 118 | 
| [jltorresm/otpgo](https://github.com/jltorresm/otpgo) | Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go. | 30 | 
| [o1egl/paseto](https://github.com/o1egl/paseto) | Platform-Agnostic Security Tokens implementation in GO (Golang) | 584 | 
| [xyproto/permissions2](https://github.com/xyproto/permissions2) |   :closed_lock_with_key: Middleware for keeping track of users, login states and permissions | 448 | 
| [zpatrick/rbac](https://github.com/zpatrick/rbac) | Minimalistic RBAC package for Go applications | 94 | 
| [ThundR67/scope](https://github.com/ThundR67/scope) | Easily Manage OAuth2 Scopes In Go | 19 | 
| [alexedwards/scs](https://github.com/alexedwards/scs) | HTTP Session Management for Go | 1063 | 
| [chmike/securecookie](https://github.com/chmike/securecookie) | Fast, secure and efficient secure cookie encoder/decoder  | 55 | 
| [icza/session](https://github.com/icza/session) | Go session management for web servers (including support for Google App Engine - GAE). | 107 | 
| [f0rmiga/sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | A driver for the SessionGate Redis module - easy session management using the Go language. | 10 | 
| [adam-hanna/sessions](https://github.com/adam-hanna/sessions) | A dead simple, highly performant, highly customizable sessions middleware for go http servers. | 63 | 
| [swithek/sessionup](https://github.com/swithek/sessionup) | Straightforward HTTP session management | 114 | 
| [brianvoe/sjwt](https://github.com/brianvoe/sjwt) | Simple JWT Golang | 97 | 


# Blockchain

| Repository | Description | Stars |
|:---|:---|:---:|
| [cosmos/cosmos-sdk](https://github.com/cosmos/cosmos-sdk) | :chains: A Framework for Building High Value Public Blockchains :sparkles: | 3697 | 
| [ethereum/go-ethereum](https://github.com/ethereum/go-ethereum) | Official Go implementation of the Ethereum protocol | 36274 | 
| [ChainSafe/gossamer](https://github.com/ChainSafe/gossamer) | 🕸️ Gossamer: A Go implementation of the Polkadot Host | 312 | 
| [gagliardetto/solana-go](https://github.com/gagliardetto/solana-go) | Go SDK library for the Solana Blockchain | 241 | 
| [tendermint/tendermint](https://github.com/tendermint/tendermint) | ⟁ Tendermint Core (BFT Consensus) in Go | 4752 | 


# Bot Building

| Repository | Description | Stars |
|:---|:---|:---:|
| [NicoNex/echotron](https://github.com/NicoNex/echotron) | An elegant and concurrent library for Telegram bots in Go. | 88 | 
| [ewohltman/ephemeral-roles](https://github.com/ewohltman/ephemeral-roles) | A Discord bot for managing ephemeral roles based upon voice channel member presence. | 57 | 
| [go-chat-bot/bot](https://github.com/go-chat-bot/bot) | IRC, Slack, Telegram and RocketChat bot written in go | 730 | 
| [go-joe - A general-purpose bot library inspired by Hubot but written in Go.](https://joe-bot.net?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [oklahomer/go-sarah](https://github.com/oklahomer/go-sarah) | Simple yet customizable bot framework written in Go. | 228 | 
| [olebedev/go-tgbot](https://github.com/olebedev/go-tgbot) | Golang  telegram bot API wrapper, session-based router and middleware | 112 | 
| [gempir/go-twitch-irc](https://github.com/gempir/go-twitch-irc) | go irc client for twitch.tv | 222 | 
| [saniales/golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | A golang implementation of a console-based trading bot for cryptocurrency exchanges | 728 | 
| [nikepan/govkbot](https://github.com/nikepan/govkbot) | VK bot package for Go | 38 | 
| [sbstjn/hanu](https://github.com/sbstjn/hanu) | Golang Framework for writing Slack bots | 136 | 
| [stellar/kelp](https://github.com/stellar/kelp) | Kelp is a free and open-source trading bot for the Stellar DEX and 100+ centralized exchanges | 868 | 
| [ezeoleaf/larry](https://github.com/ezeoleaf/larry) | Larry 🐦 is a really simple Twitter bot generator that tweets random repositories from Github built in Go | 40 | 
| [zhulik/margelet](https://github.com/zhulik/margelet) | Telegram Bot Framework for Go | 68 | 
| [onrik/micha](https://github.com/onrik/micha) | Client lib for Telegram bot api | 18 | 
| [olivia-ai/olivia](https://github.com/olivia-ai/olivia) | 💁‍♀️Your new best friend powered by an artificial neural network | 3209 | 
| [innogames/slack-bot](https://github.com/innogames/slack-bot) | Ready to use Slack bot for lazy developers: start Jenkins jobs, watch Jira tickets, watch pull requests... | 86 | 
| [shomali11/slacker](https://github.com/shomali11/slacker) | Slack Bot Framework | 588 | 
| [alexandre-normand/slackscot](https://github.com/alexandre-normand/slackscot) | Slack bot core/framework written in Go with support for reactions to message updates/deletes | 50 | 
| [yanzay/tbot](https://github.com/yanzay/tbot) | Go library for Telegram Bot API | 323 | 
| [tucnak/telebot](https://github.com/tucnak/telebot) | Telebot is a Telegram bot framework in Go. | 2449 | 
| [mymmrac/telego](https://github.com/mymmrac/telego) | Telegram Bot API library for Golang | 21 | 
| [go-telegram-bot-api/telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) | Golang bindings for the Telegram Bot API | 3495 | 
| [alfiankan/teleterm](https://github.com/alfiankan/teleterm) | Telegram Bot Exec Terminal Command | 7 | 
| [kyleterry/tenyks](https://github.com/kyleterry/tenyks) | The Tenyks IRC bot. | 172 | 


# Build Automation

| Repository | Description | Stars |
|:---|:---|:---:|
| [gopinath-langote/1build](https://github.com/gopinath-langote/1build) | Frictionless way of managing project-specific commands | 146 | 
| [GuilhermeCaruso/anko](https://github.com/GuilhermeCaruso/anko) | :crystal_ball: Simple application watcher | 22 | 
| [maxcnunes/gaper](https://github.com/maxcnunes/gaper) | Builds and restarts a Go project when it crashes or some watched file changes | 53 | 
| [gilbert - Build system and task runner for Go projects.](https://go-gilbert.github.io?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [goyek/goyek](https://github.com/goyek/goyek) | Create build pipelines in Go  | 282 | 
| [magefile/mage](https://github.com/magefile/mage) | a Make/rake-like dev tool using Go | 2920 | 
| [tj/mmake](https://github.com/tj/mmake) | Modern Make  | 1601 | 
| [oxequa/realize](https://github.com/oxequa/realize) | Realize is the #1 Golang Task Runner which enhance your workflow by automating the most common tasks and using the best performing Golang live reloading. | 4224 | 
| [go-task/task](https://github.com/go-task/task) | A task runner / simpler Make alternative written in Go | 4839 | 
| [taskctl/taskctl](https://github.com/taskctl/taskctl) | Concurrent task runner, developer's routine tasks automation toolkit. Simple modern alternative to GNU Make 🧰 | 163 | 


# Command Line

| Repository | Description | Stars |
|:---|:---|:---:|


# Advanced Console UIs

| Repository | Description | Stars |
|:---|:---|:---:|
| [guptarohit/asciigraph](https://github.com/guptarohit/asciigraph) | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. | 1884 | 
| [logrusorgru/aurora](https://github.com/logrusorgru/aurora) | Golang ultimate ANSI-colors that supports Printf/Sprintf methods | 1178 | 
| [Delta456/box-cli-maker](https://github.com/Delta456/box-cli-maker) | Make Highly Customized Boxes for your CLI | 194 | 
| [mingrammer/cfmt](https://github.com/mingrammer/cfmt) | :art: Contextual fmt inspired by bootstrap color classes | 82 | 
| [i582/cfmt](https://github.com/i582/cfmt) | Small library for simple and convenient formatted stylized output to the console. | 37 | 
| [ttacon/chalk](https://github.com/ttacon/chalk) | Intuitive package for prettifying terminal/console output. http://godoc.org/github.com/ttacon/chalk | 396 | 
| [TreyBastian/colourize](https://github.com/TreyBastian/colourize) | An ANSI colour terminal package for Go | 25 | 
| [wzshiming/ctc](https://github.com/wzshiming/ctc) | Console Text Colors - The non-invasive cross-platform terminal color library does not need to modify the Print method | 37 | 
| [workanator/go-ataman](https://github.com/workanator/go-ataman) | Another Text Attribute Manupulator | 11 | 
| [mattn/go-colorable](https://github.com/mattn/go-colorable) |  | 595 | 
| [daviddengcn/go-colortext](https://github.com/daviddengcn/go-colortext) | Change the color of console text. | 210 | 
| [mattn/go-isatty](https://github.com/mattn/go-isatty) |  | 600 | 
| [c-bata/go-prompt](https://github.com/c-bata/go-prompt) | Building powerful interactive prompts in Go, inspired by python-prompt-toolkit. | 4373 | 
| [jroimartin/gocui](https://github.com/jroimartin/gocui) | Minimalist Go package aimed at creating Console User Interfaces. | 8023 | 
| [labstack/gommon](https://github.com/labstack/gommon) | Common packages for Go | 435 | 
| [gookit/color](https://github.com/gookit/color) | 🎨 Terminal color rendering library, support 8/16 colors, 256 colors, RGB color rendering output, support Print/Sprintf methods, compatible with Windows. GO CLI 控制台颜色渲染工具库，支持16色，256色，RGB色彩渲染输出，使用类似于 Print/Sprintf，兼容并支持 Windows 环境的色彩渲染 | 1053 | 
| [cyucelen/marker](https://github.com/cyucelen/marker) |  🖍️ Marker is the easiest way to match and mark strings for colorful terminal outputs! | 25 | 
| [vbauerster/mpb](https://github.com/vbauerster/mpb) | multi progress bar for Go cli applications | 1628 | 
| [schollz/progressbar](https://github.com/schollz/progressbar) | A really basic thread-safe progress bar for Golang applications | 2368 | 
| [pterm/pterm](https://github.com/pterm/pterm) | ✨ #PTerm is a modern Go module to beautify console output. Featuring charts, progressbars, tables, trees, and much more 🚀 It's completely configurable and 100% cross-platform compatible. | 2520 | 
| [alexeyco/simpletable](https://github.com/alexeyco/simpletable) | Simple tables in terminal with Go | 350 | 
| [briandowns/spinner](https://github.com/briandowns/spinner) | Go (golang) package with 90 configurable terminal spinner/progress indicators. | 1724 | 
| [cheynewallace/tabby](https://github.com/cheynewallace/tabby) | A tiny library for super simple Golang tables | 312 | 
| [tomlazar/table](https://github.com/tomlazar/table) | pretty colorfull tables in go with less effort | 17 | 
| [InVisionApp/tabular](https://github.com/InVisionApp/tabular) | Tabular simplifies printing ASCII tables from command line utilities | 61 | 
| [nsf/termbox-go](https://github.com/nsf/termbox-go) | Pure Go termbox implementation | 4267 | 
| [mum4k/termdash](https://github.com/mum4k/termdash) | Terminal based dashboard. | 1916 | 
| [muesli/termenv](https://github.com/muesli/termenv) | Advanced ANSI style & color support for your terminal applications | 980 | 
| [gizak/termui](https://github.com/gizak/termui) | Golang terminal dashboard | 11662 | 
| [gosuri/uilive](https://github.com/gosuri/uilive) | uilive is a go library for updating terminal output in realtime | 1434 | 
| [gosuri/uiprogress](https://github.com/gosuri/uiprogress) | A go library to render progress bars in terminal applications | 1901 | 
| [gosuri/uitable](https://github.com/gosuri/uitable) | A go library to improve readability in terminal apps using tabular data | 639 | 
| [theckman/yacspin](https://github.com/theckman/yacspin) | Yet Another CLi Spinner; providing over 80 easy to use and customizable terminal spinners for multiple OSes | 324 | 


# Standard CLI

| Repository | Description | Stars |
|:---|:---|:---:|
| [cristalhq/acmd](https://github.com/cristalhq/acmd) | Simple, useful and opinionated CLI package in Go. | 47 | 
| [akamensky/argparse](https://github.com/akamensky/argparse) | Argparse for golang. Just because `flag` sucks | 409 | 
| [cosiner/argv](https://github.com/cosiner/argv) |  | 32 | 
| [rsteube/carapace](https://github.com/rsteube/carapace) | command argument completion generator for spf13/cobra | 33 | 
| [rsteube/carapace-bin](https://github.com/rsteube/carapace-bin) | multi-shell multi-command argument completer | 40 | 
| [mkideal/cli](https://github.com/mkideal/cli) | CLI - A package for building command line app with go | 641 | 
| [teris-io/cli](https://github.com/teris-io/cli) | Simple and complete API for building command line applications in Go | 106 | 
| [tucnak/climax](https://github.com/tucnak/climax) | Climax is an alternative CLI with the human face | 194 | 
| [leaanthony/clir](https://github.com/leaanthony/clir) | A Simple and Clear CLI library. Dependency free. | 87 | 
| [posener/cmd](https://github.com/posener/cmd) | The standard library flag package with its missing features | 33 | 
| [hedzr/cmdr](https://github.com/hedzr/cmdr) | POSIX-compliant command-line UI (CLI) parser and Hierarchical-configuration operations | 99 | 
| [spf13/cobra](https://github.com/spf13/cobra) | A Commander for modern Go CLI interactions | 25782 | 
| [rainu/go-command-chain](https://github.com/rainu/go-command-chain) | A go library for easy configure and run command chains. Such like pipelining in unix shells. | 23 | 
| [jaffee/commandeer](https://github.com/jaffee/commandeer) | Automatically sets up command line flags based on struct fields and tags. | 154 | 
| [posener/complete](https://github.com/posener/complete) | bash completion written in go + bash completion for go command | 823 | 
| [dnote/dnote](https://github.com/dnote/dnote) | A simple command line notebook for programmers | 2251 | 
| [elves/elvish](https://github.com/elves/elvish) | Elvish = Expressive Programming Language + Versatile Interactive Shell | 4671 | 
| [codingconcepts/env](https://github.com/codingconcepts/env) | Tag-based environment configuration for structs | 90 | 
| [cosiner/flag](https://github.com/cosiner/flag) | Flag is a simple but powerful command line option parsing library for Go support infinite level subcommand | 118 | 
| [integrii/flaggy](https://github.com/integrii/flaggy) | Idiomatic Go input parsing with subcommands, positional values, and flags at any position. No required project or package layout and no external dependencies. | 782 | 
| [sgreben/flagvar](https://github.com/sgreben/flagvar) | A collection of CLI argument types for the Go `flag` package. | 37 | 
| [RijulGulati/go-andotp](https://github.com/RijulGulati/go-andotp) | CLI program to encrypt/decrypt andOTP files | 16 | 
| [alexflint/go-arg](https://github.com/alexflint/go-arg) | Struct-based argument parsing in Go | 1366 | 
| [yitsushi/go-commander](https://github.com/yitsushi/go-commander) | Go library to simplify CLI workflow | 26 | 
| [jessevdk/go-flags](https://github.com/jessevdk/go-flags) | go command line option parser | 2186 | 
| [DavidGamba/go-getoptions](https://github.com/DavidGamba/go-getoptions) | Fully featured Go (golang) command line option parser with built-in auto-completion support. | 42 | 
| [devfacet/gocmd](https://github.com/devfacet/gocmd) | A Go library for building command line applications. | 56 | 
| [hidevopsio/hiboot](https://github.com/hidevopsio/hiboot) | hiboot is a high performance web and cli application framework with dependency injection support | 164 | 
| [liujianping/job](https://github.com/liujianping/job) | JOB, make your short-term command as a long-term job. 将命令行规划成任务的工具 | 112 | 
| [alecthomas/kingpin](https://github.com/alecthomas/kingpin) | CONTRIBUTIONS ONLY: A Go (golang) command line and flag parser | 3201 | 
| [peterh/liner](https://github.com/peterh/liner) | Pure Go line editor with history, inspired by linenoise | 883 | 
| [mitchellh/cli](https://github.com/mitchellh/cli) | A Go library for implementing command-line interfaces. | 1489 | 
| [jawher/mow.cli](https://github.com/jawher/mow.cli) | A versatile library for building CLI applications in Go | 790 | 
| [nanovms/ops](https://github.com/nanovms/ops) | ops - build and run nanos unikernels | 941 | 
| [spf13/pflag](https://github.com/spf13/pflag) | Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. | 1731 | 
| [Zaba505/sand](https://github.com/Zaba505/sand) | Package for creating interpreters | 15 | 
| [octago/sflags](https://github.com/octago/sflags) | Generate flags by parsing structures | 132 | 
| [antham/strumt](https://github.com/antham/strumt) | Strumt is a library to create prompt chain | 44 | 
| [bobg/subcmd](https://github.com/bobg/subcmd) |  | 1 | 
| [liujianping/ts](https://github.com/liujianping/ts) | timestamp convert & compare tool. 时间戳转换与对比工具 | 13 | 
| [ukautz/clif](https://github.com/ukautz/clif) | Another CLI framework for Go. It works on my machine. | 111 | 
| [urfave/cli](https://github.com/urfave/cli) | A simple, fast, and fun package for building command line apps in Go | 17575 | 
| [dixonwille/wlog](https://github.com/dixonwille/wlog) | A simple logging interface that supports cross-platform color and concurrency. | 55 | 
| [dixonwille/wmenu](https://github.com/dixonwille/wmenu) | An easy to use menu structure for cli applications that prompts users to make choices. | 153 | 


# Configuration

| Repository | Description | Stars |
|:---|:---|:---:|
| [cristalhq/aconfig](https://github.com/cristalhq/aconfig) | Simple, useful and opinionated config loader. | 355 | 
| [ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv) | ✨Clean and minimalistic environment configuration reader for Golang | 496 | 
| [golobby/config](https://github.com/golobby/config) | A lightweight yet powerful configuration manager for Go projects | 263 | 
| [JeremyLoy/config](https://github.com/JeremyLoy/config) | 12 factor configuration as a typesafe struct in as little as two function calls | 298 | 
| [olebedev/config](https://github.com/olebedev/config) | JSON or YAML configuration wrapper with convenient access methods. | 244 | 
| [BoRuDar/configuration](https://github.com/BoRuDar/configuration) | Library for setting values to structs' fields from env, flags, files or default tag | 50 | 
| [paked/configure](https://github.com/paked/configure) | Configure is a Go package that gives you easy configuration of your project through redundancy | 56 | 
| [sherifabdlnaby/configuro](https://github.com/sherifabdlnaby/configuro) | An opinionated configuration loading framework for Containerized and Cloud-Native applications. | 81 | 
| [heetch/confita](https://github.com/heetch/confita) | Load configuration in cascade from multiple backends into a struct | 431 | 
| [the4thamigo-uk/conflate](https://github.com/the4thamigo-uk/conflate) | Library providing routines to merge and validate JSON, YAML and/or TOML files | 24 | 
| [caarlos0/env](https://github.com/caarlos0/env) | A simple and zero-dependencies library to parse environment variables into structs. | 2343 | 
| [junk1tm/env](https://github.com/junk1tm/env) | A lightweight package for loading environment variables into structs | 14 | 
| [tomazk/envcfg](https://github.com/tomazk/envcfg) | Un-marshaling environment variables to Go structs | 97 | 
| [ian-kent/envconf](https://github.com/ian-kent/envconf) | Configure Go applications from the environment | 10 | 
| [vrischmann/envconfig](https://github.com/vrischmann/envconfig) | Small library to read your configuration from environment variables | 221 | 
| [antham/envh](https://github.com/antham/envh) | Go helpers to manage environment variables | 95 | 
| [kkyr/fig](https://github.com/kkyr/fig) | A minimalist Go configuration library | 195 | 
| [go-gcfg/gcfg](https://github.com/go-gcfg/gcfg) | read INI-style configuration files into Go structs; supports user-defined types and subsections | 159 | 
| [sakirsensoy/genv](https://github.com/sakirsensoy/genv) | Genv is a library for Go (golang) that makes it easy to read and use environment variables in your projects. It also allows environment variables to be loaded from the .env file. | 28 | 
| [PaddleHQ/go-aws-ssm](https://github.com/PaddleHQ/go-aws-ssm) | Go package that interfaces with AWS System Manager | 47 | 
| [ThomasObenaus/go-conf](https://github.com/ThomasObenaus/go-conf) | Library for easy configuration of a golang service | 2 | 
| [subpop/go-ini](https://github.com/subpop/go-ini) | automatic mirror of https://git.sr.ht/~spc/go-ini | 7 | 
| [ianlopshire/go-ssm-config](https://github.com/ianlopshire/go-ssm-config) | Go utility for loading configuration parameters from AWS SSM (Parameter Store) | 11 | 
| [ufoscout/go-up](https://github.com/ufoscout/go-up) | go-up! A simple configuration library with recursive placeholders resolution and no magic. | 37 | 
| [gosidekick/goconfig](https://github.com/gosidekick/goconfig) | goconfig uses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. | 152 | 
| [joho/godotenv](https://github.com/joho/godotenv) | A Go port of Ruby's dotenv library (Loads environment variables from `.env`.) | 4716 | 
| [ian-kent/gofigure](https://github.com/ian-kent/gofigure) | Go configuration made easy! | 62 | 
| [One-com/gone](https://github.com/One-com/gone) | Golang packages for writing small daemons and servers. | 40 | 
| [miladabc/gonfig](https://github.com/miladabc/gonfig) | Tag based configuration loader from different providers | 3 | 
| [gookit/config](https://github.com/gookit/config) | 📝 Go config manage(load,get,set). support JSON, YAML, TOML, INI, HCL, ENV and Flags. Multi file load, data override merge, parse ENV var. Go应用配置加载管理，支持多种格式，多文件加载，远程文件加载，支持数据合并，解析环境变量名 | 335 | 
| [beatlabs/harvester](https://github.com/beatlabs/harvester) | Harvest configuration, watch and notify subscriber | 98 | 
| [hjson/hjson-go](https://github.com/hjson/hjson-go) | Hjson for Go | 260 | 
| [gurkankaymak/hocon](https://github.com/gurkankaymak/hocon) | go implementation of lightbend's HOCON configuration library https://github.com/lightbend/config | 38 | 
| [schachmat/ingo](https://github.com/schachmat/ingo) | persistent storage for flags in go | 36 | 
| [go-ini/ini](https://github.com/go-ini/ini) | Package ini provides INI file read and write functionality in Go | 2900 | 
| [wlevene/ini](https://github.com/wlevene/ini) | ini parser for golang | 8 | 
| [joshbetz/config](https://github.com/joshbetz/config) | 🛠 A configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. | 208 | 
| [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) | Golang library for managing configuration data from environment variables | 4034 | 
| [knadh/koanf](https://github.com/knadh/koanf) | Simple, lightweight, extensible, configuration management library for Go. Support for JSON, TOML, YAML, env, command line, file, S3 etc. Alternative to viper. | 916 | 
| [lalamove/konfig](https://github.com/lalamove/konfig) | Composable, observable and performant config handling for Go for the distributed processing era | 625 | 
| [alecthomas/kong](https://github.com/alecthomas/kong) | Kong is a command-line parser for Go | 889 | 
| [sasbury/mini](https://github.com/sasbury/mini) | A golang package for parsing ini-style configuration files | 29 | 
| [nasermirzaei89/env](https://github.com/nasermirzaei89/env) | Golang Get Environment Variables Package | 8 | 
| [goraz/onion](https://github.com/goraz/onion) | Layer based configuration for golang | 96 | 
| [Yiling-J/piper](https://github.com/Yiling-J/piper) | 🛠 Viper wrapper with config inheritance and key generation | 4 | 
| [tucnak/store](https://github.com/tucnak/store) | A dead simple configuration manager for Go applications | 259 | 
| [oblq/swap](https://github.com/oblq/swap) | Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env). | 5 | 
| [diegomarangoni/typenv](https://github.com/diegomarangoni/typenv) | Go minimalist typed environment variables library | 5 | 
| [omeid/uconfig](https://github.com/omeid/uconfig) | Lightweight, zero-dependency, and extendable configuration management library for Go | 42 | 
| [spf13/viper](https://github.com/spf13/viper) | Go configuration with fangs | 18671 | 
| [adrg/xdg](https://github.com/adrg/xdg) | Go implementation of the XDG Base Directory Specification and XDG user directories | 191 | 
| [OpenPeeDeeP/xdg](https://github.com/OpenPeeDeeP/xdg) | A cross platform package that follows the XDG Standard | 66 | 


# Continuous Integration

| Repository | Description | Stars |
|:---|:---|:---:|
| [ovh/cds](https://github.com/ovh/cds) | Enterprise-Grade Continuous Delivery & DevOps Automation Open Source Platform | 3782 | 
| [harness/drone](https://github.com/harness/drone) | Drone is a Container-Native, Continuous Delivery Platform | 24709 | 
| [duck8823/duci](https://github.com/duck8823/duci) | The simple ci server  | 70 | 
| [nikogura/gomason](https://github.com/nikogura/gomason) | A tool for testing, building, signing, and publishing binaries. | 52 | 
| [haveyoudebuggedit/gotestfmt](https://github.com/haveyoudebuggedit/gotestfmt) | go test output for humans | 201 | 
| [mattn/goveralls](https://github.com/mattn/goveralls) |  | 714 | 
| [go-playground/overalls](https://github.com/go-playground/overalls) | :jeans:Multi-Package go project coverprofile for tools like goveralls | 109 | 
| [lawrencewoodman/roveralls](https://github.com/lawrencewoodman/roveralls) | A Go recursive coverage testing tool | 16 | 


# CSS Preprocessors

| Repository | Description | Stars |
|:---|:---|:---:|
| [yosssi/gcss](https://github.com/yosssi/gcss) | Pure Go CSS Preprocessor | 448 | 
| [wellington/go-libsass](https://github.com/wellington/go-libsass) | Go wrapper for libsass, the only Sass 3.5 compiler for Go | 185 | 


# Data Structures

| Repository | Description | Stars |
|:---|:---|:---:|
| [shady831213/algorithms](https://github.com/shady831213/algorithms) | CLRS study. Codes are written with golang. | 626 | 
| [iancmcc/bingo](https://github.com/iancmcc/bingo) | Pack native Golang types to bytes while preserving sort order. Perfect for key-value stores or binary trees! | 3 | 
| [zhuangsirui/binpacker](https://github.com/zhuangsirui/binpacker) | A binary stream packer and unpacker | 183 | 
| [yourbasic/bit](https://github.com/yourbasic/bit) | Bitset data structure | 114 | 
| [kelindar/bitmap](https://github.com/kelindar/bitmap) | Simple dense bitmap index in Go with binary operators | 128 | 
| [bits-and-blooms/bitset](https://github.com/bits-and-blooms/bitset) | Go package implementing bitsets | 853 | 
| [bits-and-blooms/bloom](https://github.com/bits-and-blooms/bloom) | Go package implementing Bloom filters | 1465 | 
| [zentures/bloom](https://github.com/zentures/bloom) | Bloom filters implemented in Go. | 146 | 
| [yourbasic/bloom](https://github.com/yourbasic/bloom) | Probabilistic set data structure | 69 | 
| [OldPanda/bloomfilter](https://github.com/OldPanda/bloomfilter) | Yet another Bloomfilter implementation in Go, compatible with Java's Guava library | 9 | 
| [tylertreat/BoomFilters](https://github.com/tylertreat/BoomFilters) | Probabilistic data structures for processing continuous, unbounded streams. | 1416 | 
| [lrita/cmap](https://github.com/lrita/cmap) | a thread-safe concurrent map for go | 24 | 
| [free/concurrent-writer](https://github.com/free/concurrent-writer) | Highly concurrent drop-in replacement for bufio.Writer | 39 | 
| [InVisionApp/conjungo](https://github.com/InVisionApp/conjungo) | A small flexible merge library in go | 103 | 
| [seiflotfy/count-min-log](https://github.com/seiflotfy/count-min-log) | Go implementation of Count-Min-Log | 55 | 
| [superwhiskers/crunch](https://github.com/superwhiskers/crunch) | take bytes out of things easily ✨🍪 | 55 | 
| [linvon/cuckoo-filter](https://github.com/linvon/cuckoo-filter) | Cuckoo Filter go implement, better than Bloom Filter, configurable and space optimized  布谷鸟过滤器的Go实现，优于布隆过滤器，可以定制化过滤器参数，并进行了空间优化 | 203 | 
| [seiflotfy/cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | Cuckoo Filter: Practically Better Than Bloom | 889 | 
| [edwingeng/deque](https://github.com/edwingeng/deque) | A highly optimized double-ended queue | 37 | 
| [gammazero/deque](https://github.com/gammazero/deque) | Fast ring-buffer deque (double-ended queue) | 310 | 
| [srfrog/dict](https://github.com/srfrog/dict) | Python-like dictionaries for Go | 23 | 
| [ihebu/dsu](https://github.com/ihebu/dsu) | Disjoint Set data structure implementation in Go | 6 | 
| [zentures/encoding](https://github.com/zentures/encoding) | Integer Compression Libraries for Go | 118 | 
| [cocoonspace/fsm](https://github.com/cocoonspace/fsm) | Finite State Machine package in Go | 15 | 
| [ulovecode/gdcache](https://github.com/ulovecode/gdcache) | gdcache is a pure non-intrusive cache library implemented by golang, you can use it to implement your own cache. | 9 | 
| [plar/go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | Adaptive Radix Trees implemented in Go | 228 | 
| [Workiva/go-datastructures](https://github.com/Workiva/go-datastructures) | A collection of useful, performant, and threadsafe Go datastructures. | 6426 | 
| [hbollon/go-edlib](https://github.com/hbollon/go-edlib) | 📚 String comparison and edit distance algorithms library, featuring : Levenshtein, LCS, Hamming, Damerau levenshtein (OSA and Adjacent transpositions algorithms), Jaro-Winkler, Cosine, etc... | 314 | 
| [amallia/go-ef](https://github.com/amallia/go-ef) | A Go implementation of the Elias-Fano encoding | 19 | 
| [hailocab/go-geoindex](https://github.com/hailocab/go-geoindex) | Go native library for fast point tracking and K-Nearest queries | 338 | 
| [OrlovEvgeny/go-mcache](https://github.com/OrlovEvgeny/go-mcache) | Fast in-memory key:value store/cache with TTL | 83 | 
| [arl/go-rquad](https://github.com/arl/go-rquad) | :pushpin: State of the art point location and neighbour finding algorithms for region quadtrees, in Go | 120 | 
| [eko/gocache](https://github.com/eko/gocache) | ☔️ A complete Go cache library that brings you multiple ways of managing your caches | 1117 | 
| [enriquebris/goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | Go concurrent-safe, goroutine-safe, thread-safe queue | 167 | 
| [emirpasic/gods](https://github.com/emirpasic/gods) | GoDS (Go Data Structures). Containers (Sets, Lists, Stacks, Maps, Trees), Sets (HashSet, TreeSet, LinkedHashSet), Lists (ArrayList, SinglyLinkedList, DoublyLinkedList), Stacks (LinkedListStack, ArrayStack), Maps (HashMap, TreeMap, HashBidiMap, TreeBidiMap, LinkedHashMap), Trees (RedBlackTree, AVLTree, BTree, BinaryHeap), Comparators, Iterators, Enumerables, Sort, JSON | 11210 | 
| [xxjwxc/gofal](https://github.com/xxjwxc/gofal) | fractional api base on golang . golang math tools fractional molecular denominator 分数计算 分子 分母 运算 | 13 | 
| [deckarep/golang-set](https://github.com/deckarep/golang-set) | A simple set type for the Go language. Trusted by Docker, 1Password, Ethereum and Hashicorp. | 2299 | 
| [zoumo/goset](https://github.com/zoumo/goset) | Set is a useful collection but there is no built-in implementation in Go lang. | 46 | 
| [ryszard/goskiplist](https://github.com/ryszard/goskiplist) | A skip list implementation in Go | 231 | 
| [liyue201/gostl](https://github.com/liyue201/gostl) | Data structure and algorithm library for go, designed to provide functions similar to C++ STL | 626 | 
| [go-gota/gota](https://github.com/go-gota/gota) | Gota: DataFrames and data wrangling in Go (Golang) | 2028 | 
| [yaa110/goterator](https://github.com/yaa110/goterator) | Lazy iterator implementation for Golang | 6 | 
| [bobg/hashsplit](https://github.com/bobg/hashsplit) |  | 7 | 
| [emvi/hide](https://github.com/emvi/hide) | ID type with marshalling to/from hash to prevent sending IDs to clients. | 43 | 
| [google/hilbert](https://github.com/google/hilbert) | Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. | 246 | 
| [axiomhq/hyperloglog](https://github.com/axiomhq/hyperloglog) | HyperLogLog with lots of sugar (Sparse, LogLog-Beta bias correction and TailCut space reduction) | 772 | 
| [disksing/iter](https://github.com/disksing/iter) | Go implementation of C++ STL iterators and algorithms. | 153 | 
| [agext/levenshtein](https://github.com/agext/levenshtein) | Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. | 67 | 
| [agnivade/levenshtein](https://github.com/agnivade/levenshtein) | Go implementation to calculate Levenshtein Distance. | 185 | 
| [embano1/memlog](https://github.com/embano1/memlog) | A Kafka log inspired in-memory and append-only data structure | 38 | 
| [bobg/merkle](https://github.com/bobg/merkle) | Merkle hash trees | 2 | 
| [cbergoon/merkletree](https://github.com/cbergoon/merkletree) | A Merkle Tree implementation written in Go. | 321 | 
| [BlackRabbitt/mspm](https://github.com/BlackRabbitt/mspm) | Multi-String Pattern Matching Algorithm Using TrieHashNode | 17 | 
| [kak-tus/nan](https://github.com/kak-tus/nan) | Zero allocation Nullable structures in one library with handy conversion functions, marshallers and unmarshallers | 48 | 
| [emvi/null](https://github.com/emvi/null) | Nullable Go types that can be marshalled/unmarshalled to/from JSON. | 24 | 
| [tejzpr/ordered-concurrently](https://github.com/tejzpr/ordered-concurrently) | Ordered-concurrently a library for concurrent processing with ordered output in Go. Process work concurrently and returns output in a channel in the order of input. It is useful in concurrently processing items in a queue, and get output in the order provided by the queue. | 13 | 
| [nazar256/parapipe](https://github.com/nazar256/parapipe) | Paralleling pipeline | 16 | 
| [MonaxGT/parsefields](https://github.com/MonaxGT/parsefields) | Tools for parse JSON-like logs for collecting unique fields and events | 6 | 
| [hyfather/pipeline](https://github.com/hyfather/pipeline) | Pipelines using goroutines | 40 | 
| [viant/ptrie](https://github.com/viant/ptrie) | A prefix tree implementation in go  | 19 | 
| [rocketlaunchr/remember-go](https://github.com/rocketlaunchr/remember-go) | Cache Slow Database Queries | 109 | 
| [tannerryan/ring](https://github.com/tannerryan/ring) | Package ring provides a high performance and thread safe Go implementation of a bloom filter. | 124 | 
| [RoaringBitmap/roaring](https://github.com/RoaringBitmap/roaring) | Roaring bitmaps in Go (golang) | 1513 | 
| [StudioSol/set](https://github.com/StudioSol/set) | A simple Set data structure implementation in Go (Golang) using LinkedHashMap. | 19 | 
| [MauriceGit/skiplist](https://github.com/MauriceGit/skiplist) | A Go library for an efficient implementation of a skip list: https://godoc.org/github.com/MauriceGit/skiplist | 194 | 
| [gansidui/skiplist](https://github.com/gansidui/skiplist) | skiplist for golang | 77 | 
| [srfrog/slices](https://github.com/srfrog/slices) | Functions that operate on slices. Similar to functions from package strings or package bytes that have been adapted to work with slices. | 7 | 
| [zekroTJA/timedmap](https://github.com/zekroTJA/timedmap) | A thread safe map which has expiring key-value pairs. | 41 | 
| [perdata/treap](https://github.com/perdata/treap) | golang persistent immutable treap sorted sets | 15 | 
| [igrmk/treemap](https://github.com/igrmk/treemap) | Generic sorted map for Go with red-black tree under the hood | 16 | 
| [derekparker/trie](https://github.com/derekparker/trie) | Data structure and relevant algorithms for extremely fast prefix/fuzzy string searching. | 576 | 
| [jellydator/ttlcache](https://github.com/jellydator/ttlcache) | An in-memory cache with item expiration and generics | 359 | 
| [gurukami/typ](https://github.com/gurukami/typ) | Null Types, Safe primitive type conversion and fetching value from complex structures. | 30 | 


# Database

| Repository | Description | Stars |
|:---|:---|:---:|
| [dgraph-io/badger](https://github.com/dgraph-io/badger) | Fast key-value DB in Go. | 10561 | 
| [etcd-io/bbolt](https://github.com/etcd-io/bbolt) | An embedded key/value database for Go. | 5369 | 
| [iwanbk/bcache](https://github.com/iwanbk/bcache) | Eventually consistent distributed in-memory  cache Go library | 83 | 
| [allegro/bigcache](https://github.com/allegro/bigcache) | Efficient cache for gigabytes of data written in Go. | 5529 | 
| [Bitcask - Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL).](https://git.mills.io/prologic/bitcask?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [tidwall/buntdb](https://github.com/tidwall/buntdb) | BuntDB is an embeddable, in-memory key/value database for Go with custom indexing and geospatial support | 3682 | 
| [akyoto/cache](https://github.com/akyoto/cache) | :handbag: Cache arbitrary data with an expiration time. | 105 | 
| [muesli/cache2go](https://github.com/muesli/cache2go) | Concurrency-safe Go caching library with expiration capabilities and access counters | 1699 | 
| [ostafen/clover](https://github.com/ostafen/clover) | A lightweight document-oriented NoSQL database written in pure Golang. | 104 | 
| [oaStuff/clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | golang bigcache with clustering as a library. | 41 | 
| [cockroachdb/cockroach](https://github.com/cockroachdb/cockroach) | CockroachDB - the open source, cloud-native distributed SQL database. | 24043 | 
| [claygod/coffer](https://github.com/claygod/coffer) | Simply ACID* key-value database. At the medium or even low latency it tries to provide greater throughput without losing the ACID properties of the database. The database provides the ability to create record headers at own discretion and use them as transactions. The maximum size of stored data is limited by the size of the computer's RAM. | 28 | 
| [kelindar/column](https://github.com/kelindar/column) | High-performance, columnar, in-memory store with bitmap indexing in Go | 885 | 
| [codingsince1985/couchcache](https://github.com/codingsince1985/couchcache) | A RESTful caching micro-service in Go backed by Couchbase | 56 | 
| [CovenantSQL/CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) | A decentralized, trusted, high performance, SQL database with blockchain features | 1292 | 
| [securitybunker/databunker](https://github.com/securitybunker/databunker) | Secure SDK/vault for personal records/PII built to comply with GDPR | 965 | 
| [dgraph-io/dgraph](https://github.com/dgraph-io/dgraph) | Native GraphQL Database with graph backend | 17812 | 
| [peterbourgon/diskv](https://github.com/peterbourgon/diskv) | A disk-backed key-value store. | 1146 | 
| [dtm-labs/dtf](https://github.com/dtm-labs/dtf) | A distributed transaction framework that supports multiple languages, supports saga, tcc, xa, 2-phase message strategies. | 238 | 
| [krotik/eliasdb](https://github.com/krotik/eliasdb) | EliasDB a graph-based database. | 881 | 
| [VictoriaMetrics/fastcache](https://github.com/VictoriaMetrics/fastcache) | Fast thread-safe inmemory cache for big number of entries in Go. Minimizes GC overhead | 1393 | 
| [bluele/gcache](https://github.com/bluele/gcache) | An in-memory cache library for golang. It supports multiple eviction policies: LRU, LFU, ARC | 1864 | 
| [patrickmn/go-cache](https://github.com/patrickmn/go-cache) | An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. | 5982 | 
| [HDT3213/godis](https://github.com/HDT3213/godis) | A Golang implemented Redis Server and Cluster. Go 语言实现的 Redis 服务器和分布式集群 | 1713 | 
| [syndtr/goleveldb](https://github.com/syndtr/goleveldb) | LevelDB key/value database in Go. | 4985 | 
| [golang/groupcache](https://github.com/golang/groupcache) | groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. | 11256 | 
| [jameycribbs/hare](https://github.com/jameycribbs/hare) | Hare is a nimble little database management system for Go. | 53 | 
| [codenotary/immudb](https://github.com/codenotary/immudb) | immudb - immutable database based on zero trust, SQL and Key-Value, tamperproof, data change history | 7327 | 
| [influxdata/influxdb](https://github.com/influxdata/influxdb) | Scalable datastore for metrics, events, and real-time analytics | 23220 | 
| [go-kivik/kivik](https://github.com/go-kivik/kivik) | Kivik provides a common interface to CouchDB or CouchDB-like databases for Go and GopherJS. | 229 | 
| [ledisdb/ledisdb](https://github.com/ledisdb/ledisdb) | A high performance NoSQL Database Server powered by Go | 3831 | 
| [jmhodges/levigo](https://github.com/jmhodges/levigo) | levigo is a Go wrapper for LevelDB | 404 | 
| [flower-corp/lotusdb](https://github.com/flower-corp/lotusdb) | Fast k/v storage compatible with lsm tree and b+tree, inspired by SLM-DB in USENIX FAST ’19. | 330 | 
| [milvus-io/milvus](https://github.com/milvus-io/milvus) | An open-source vector database for scalable similarity search and AI applications. | 9861 | 
| [couchbase/moss](https://github.com/couchbase/moss) | moss - a simple, fast, ordered, persistable, key-val storage library for golang | 860 | 
| [nutsdb/nutsdb](https://github.com/nutsdb/nutsdb) | A simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set. | 2007 | 
| [fern4lvarez/piladb](https://github.com/fern4lvarez/piladb) | Lightweight RESTful database engine based on stack data structures | 194 | 
| [akrylysov/pogreb](https://github.com/akrylysov/pogreb) | Embedded key-value store for read-heavy workloads written in Go | 878 | 
| [prometheus/prometheus](https://github.com/prometheus/prometheus) | The Prometheus monitoring system and time series database. | 41654 | 
| [recoilme/pudge](https://github.com/recoilme/pudge) | Fast and simple key/value store written using Go's standard library | 315 | 
| [flower-corp/rosedb](https://github.com/flower-corp/rosedb) | 🚀A fast, stable and embedded k-v storage in Go, supports string, list, hash, set, and sorted set. | 2406 | 
| [rqlite/rqlite](https://github.com/rqlite/rqlite) | The lightweight, distributed relational database built on SQLite | 9806 | 
| [nanobox-io/golang-scribble](https://github.com/nanobox-io/golang-scribble) | A tiny Golang JSON database | 146 | 
| [rafaeljesus/tempdb](https://github.com/rafaeljesus/tempdb) | Key-value store for temporary items :memo: | 16 | 
| [pingcap/tidb](https://github.com/pingcap/tidb) | TiDB is an open source distributed HTAP database compatible with the MySQL protocol  | 30777 | 
| [HouzuoGuo/tiedot](https://github.com/HouzuoGuo/tiedot) | A rudimentary implementation of a basic document (NoSQL) database in Go | 2659 | 
| [cheshir/ttlcache](https://github.com/cheshir/ttlcache) | Simple in-memory key-value storage with TTL for each record. | 5 | 
| [unit-io/unitdb](https://github.com/unit-io/unitdb) | Fast specialized time-series database for IoT, real-time internet connected devices and AI analytics. | 86 | 
| [chrislusf/vasto](https://github.com/chrislusf/vasto) | A distributed key-value store. On Disk. Able to grow or shrink without service interruption. | 236 | 
| [VictoriaMetrics/VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) | VictoriaMetrics: fast, cost-effective monitoring solution and time series database | 6005 | 


# Date and Time

| Repository | Description | Stars |
|:---|:---|:---:|
| [golang-module/carbon](https://github.com/golang-module/carbon) | A simple, semantic and developer-friendly golang package for datetime | 1753 | 
| [uniplaces/carbon](https://github.com/uniplaces/carbon) | Carbon for Golang, an extension for Time | 691 | 
| [1set/cronrange](https://github.com/1set/cronrange) | time range expression in cron style | 15 | 
| [rickb777/date](https://github.com/rickb777/date) | A Go package for working with dates | 88 | 
| [araddon/dateparse](https://github.com/araddon/dateparse) | GoLang Parse many date strings without knowing format in advance. | 1653 | 
| [hako/durafmt](https://github.com/hako/durafmt) | :clock8:  Better time duration formatting in Go!  | 429 | 
| [wlbr/feiertage](https://github.com/wlbr/feiertage) | Gesetzliche Feiertage und mehr in Deutschland und Österreich (Bank holidays/public holidays in Austria and Germany) | 41 | 
| [yaa110/go-persian-calendar](https://github.com/yaa110/go-persian-calendar) | The implementation of Persian (Solar Hijri) Calendar in Go | 115 | 
| [xhit/go-str2duration](https://github.com/xhit/go-str2duration) | Convert string to duration in golang | 39 | 
| [nathan-osman/go-sunrise](https://github.com/nathan-osman/go-sunrise) | Go package for calculating the sunrise and sunset times for a given location | 42 | 
| [stoewer/go-week](https://github.com/stoewer/go-week) | A Go package to work with ISO 8601 week dates | 7 | 
| [bykof/gostradamus](https://github.com/bykof/gostradamus) | Gostradamus: Better DateTimes for Go 🕰️ | 160 | 
| [relvacode/iso8601](https://github.com/relvacode/iso8601) | A fast ISO8601 date parser for Go | 98 | 
| [GuilhermeCaruso/kair](https://github.com/GuilhermeCaruso/kair) | :clock1: Date and Time - Golang Formatting Library | 23 | 
| [jinzhu/now](https://github.com/jinzhu/now) | Now is a time toolkit for golang | 3613 | 
| [kirillDanshin/nulltime](https://github.com/kirillDanshin/nulltime) |  | 12 | 
| [awoodbeck/strftime](https://github.com/awoodbeck/strftime) | C99-compatible strftime formatter for use with Go time.Time instances. | 9 | 
| [SaidinWoT/timespan](https://github.com/SaidinWoT/timespan) | Golang package to manipulate time intervals. | 78 | 
| [leekchan/timeutil](https://github.com/leekchan/timeutil) | timeutil - useful extensions (Timedelta, Strftime, ...) to the golang's time package | 188 | 
| [osteele/tuesday](https://github.com/osteele/tuesday) | Ruby-compatible strftime for golang | 11 | 


# Distributed Systems

| Repository | Description | Stars |
|:---|:---|:---:|
| [lesismal/arpc](https://github.com/lesismal/arpc) | More effective network communication, two-way calling, notify and broadcast supported. | 486 | 
| [svcavallar/celeriac.v1](https://github.com/svcavallar/celeriac.v1) | Golang client library for adding support for interacting and monitoring Celery workers, tasks and events. | 70 | 
| [buraksezer/consistent](https://github.com/buraksezer/consistent) | Consistent hashing with bounded loads in Golang | 460 | 
| [mbrostami/consistenthash](https://github.com/mbrostami/consistenthash) | A Go library that implements Consistent Hashing | 10 | 
| [anacrolix/dht](https://github.com/anacrolix/dht) | dht is used by anacrolix/torrent, and is intended for use as a library in other projects both torrent related and otherwise | 226 | 
| [digota/digota](https://github.com/digota/digota) | ecommerce microservice | 436 | 
| [dotchain/dot](https://github.com/dotchain/dot) | distributed data sync with operational transformation/transforms  | 67 | 
| [edwingeng/doublejump](https://github.com/edwingeng/doublejump) | A revamped Google's jump consistent hash | 70 | 
| [lni/dragonboat](https://github.com/lni/dragonboat) | A feature complete and high performance multi-group Raft library in Go.   | 4150 | 
| [dgruber/drmaa](https://github.com/dgruber/drmaa) | Compute cluster (HPC) job submission library for Go (#golang) based on the open DRMAA standard. | 37 | 
| [dynamolock - DynamoDB-backed distributed locking implementation.](https://cirello.io/dynamolock?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [tylfin/dynatomic](https://github.com/tylfin/dynatomic) | Dynatomic is a library for using dynamodb as an atomic counter | 14 | 
| [emitter-io/emitter](https://github.com/emitter-io/emitter) | High performance, distributed and low latency publish-subscribe platform. | 3210 | 
| [andy2046/failured](https://github.com/andy2046/failured) | Adaptive Accrual Failure Detector | 4 | 
| [vectaport/flowgraph](https://github.com/vectaport/flowgraph) | Flowgraph package for scalable asynchronous system development | 47 | 
| [chrislusf/gleam](https://github.com/chrislusf/gleam) | Fast, efficient, and scalable distributed map/reduce system, DAG execution, in memory or on disk, written in pure Go, runs standalone or distributedly. | 3019 | 
| [chrislusf/glow](https://github.com/chrislusf/glow) | Glow is an easy-to-use distributed computation system written in Go, similar to Hadoop Map Reduce, Spark, Flink, Storm, etc. I am also working on another similar pure Go system, https://github.com/chrislusf/gleam , which is more flexible and more performant. | 3042 | 
| [gmsec/micro](https://github.com/gmsec/micro) | A Go distributed systems development framework | 18 | 
| [unionj-cloud/go-doudou](https://github.com/unionj-cloud/go-doudou) | go-doudou（doudou pronounce /dəudəu/）is a gossip protocol and OpenAPI 3.0 spec based decentralized microservice framework. It supports monolith service application as well. Currently, it supports RESTful service only. | 443 | 
| [InVisionApp/go-health](https://github.com/InVisionApp/go-health) | Library for enabling asynchronous health checks in your service | 619 | 
| [dgryski/go-jump](https://github.com/dgryski/go-jump) | go-jump: Jump consistent hashing | 339 | 
| [go-kit/kit](https://github.com/go-kit/kit) | A standard library for microservices. | 22661 | 
| [asim/go-micro](https://github.com/asim/go-micro) | A Go microservices framework | 17933 | 
| [sanketplus/go-mysql-lock](https://github.com/sanketplus/go-mysql-lock) | MySQL Backed Locking Primitive | 39 | 
| [pdupub/go-pdu](https://github.com/pdupub/go-pdu) | Parallel Digital Universe - A decentralized identity-based social network | 35 | 
| [AppsFlyer/go-sundheit](https://github.com/AppsFlyer/go-sundheit) | A library built to provide support for defining service health for golang services. It allows you to register async health checks for your dependencies and the service itself, provides a health endpoint that exposes their status, and health metrics. | 467 | 
| [zeromicro/go-zero](https://github.com/zeromicro/go-zero) | A web and RPC framework written in Go. It's born to ensure the stability of the busy sites with resilient design. Builtin goctl greatly improves the development productivity. | 15943 | 
| [valyala/gorpc](https://github.com/valyala/gorpc) | Simple, fast and scalable golang rpc library for high load | 646 | 
| [grpc/grpc-go](https://github.com/grpc/grpc-go) | The Go language implementation of gRPC. HTTP/2 based RPC | 15576 | 
| [hprose/hprose-golang](https://github.com/hprose/hprose-golang) | Hprose is a cross-language RPC. This project is Hprose for Golang. | 1201 | 
| [osamingo/jsonrpc](https://github.com/osamingo/jsonrpc) | The jsonrpc package helps implement of JSON-RPC 2.0 | 160 | 
| [ybbus/jsonrpc](https://github.com/ybbus/jsonrpc) | A simple go implementation of json rpc 2.0 client over http | 212 | 
| [go-kratos/kratos](https://github.com/go-kratos/kratos) | Your ultimate Go microservices framework for the cloud-native era. | 17090 | 
| [liftbridge-io/liftbridge](https://github.com/liftbridge-io/liftbridge) | Lightweight, fault-tolerant message streams. | 2250 | 
| [luraproject/lura](https://github.com/luraproject/lura) | Ultra performant API Gateway with middlewares. A project hosted at The Linux Foundation | 4953 | 
| [micro/micro](https://github.com/micro/micro) | API first cloud platform | 10993 | 
| [nats-io/nats-server](https://github.com/nats-io/nats-server) | High-Performance server for NATS.io, the cloud and edge native messaging system. | 10677 | 
| [italolelis/outboxer](https://github.com/italolelis/outboxer) | A library that implements the outboxer pattern in go | 68 | 
| [pglock - PostgreSQL-backed distributed locking implementation.](https://cirello.io/pglock?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [pjrpc - Golang JSON-RPC Server-Client with Protobuf spec.](https://gitlab.com/pjrpc/pjrpc?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [hashicorp/raft](https://github.com/hashicorp/raft) | Golang implementation of the Raft consensus protocol | 5763 | 
| [etcd-io/etcd](https://github.com/etcd-io/etcd) | Distributed reliable key-value store for the most critical data of a distributed system | 39238 | 
| [cenkalti/rain](https://github.com/cenkalti/rain) | 🌧 BitTorrent client and library in Go | 683 | 
| [bsm/redislock](https://github.com/bsm/redislock) | Simplified distributed locking implementation using Redis | 639 | 
| [resgate - Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.](https://resgate.io/?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [uber/ringpop-go](https://github.com/uber/ringpop-go) | Scalable, fault-tolerant application-layer sharding for Go applications | 728 | 
| [smallnest/rpcx](https://github.com/smallnest/rpcx) | Best microservices framework in Go, like alibaba Dubbo, but with more features, Scale easily. Try it. Test it. If you feel it's better, use it! 𝐉𝐚𝐯𝐚有𝐝𝐮𝐛𝐛𝐨, 𝐆𝐨𝐥𝐚𝐧𝐠有𝐫𝐩𝐜𝐱! | 6821 | 
| [jexia/semaphore](https://github.com/jexia/semaphore) | Take control of your data, connect with anything, and expose it anywhere through protocols such as HTTP, GraphQL, and gRPC. | 70 | 
| [ursiform/sleuth](https://github.com/ursiform/sleuth) | A Go library for master-less peer-to-peer autodiscovery and RPC between HTTP services | 347 | 
| [anacrolix/torrent](https://github.com/anacrolix/torrent) | Full-featured BitTorrent client package and utilities | 4262 | 


# Dynamic DNS

| Repository | Description | Stars |
|:---|:---|:---:|
| [skibish/ddns](https://github.com/skibish/ddns) | Personal DDNS client with Digital Ocean Networking DNS as backend. | 205 | 
| [dyndns - Background Go process to regularly and automatically check your IP Address and make updates to (one or many) Dynamic DNS records for Google domains whenever your address changes.](https://gitlab.com/alcastle/dyndns?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [TimothyYe/godns](https://github.com/TimothyYe/godns) | A dynamic DNS client tool supports AliDNS, Cloudflare, Google Domains, DNSPod, HE.net & DuckDNS & DreamHost, etc, written in Go. | 998 | 


# Email

| Repository | Description | Stars |
|:---|:---|:---:|
| [chasquid - SMTP server written in Go.](https://blitiri.com.ar/p/chasquid?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [aymerick/douceur](https://github.com/aymerick/douceur) | A simple CSS parser and inliner in Go | 211 | 
| [jordan-wright/email](https://github.com/jordan-wright/email) | Robust and flexible email library for Go | 1979 | 
| [AfterShip/email-verifier](https://github.com/AfterShip/email-verifier) | :white_check_mark: A Go library for email verification without sending any emails. | 386 | 
| [toorop/go-dkim](https://github.com/toorop/go-dkim) | DKIM package for golang | 75 | 
| [go-email-validator/go-email-validator](https://github.com/go-email-validator/go-email-validator) | 📧 Golang Email address validator | 23 | 
| [emersion/go-imap](https://github.com/emersion/go-imap) |  :inbox_tray: An IMAP library for clients and servers | 1502 | 
| [emersion/go-message](https://github.com/emersion/go-message) | :envelope: A streaming Go library for the Internet Message Format and mail messages | 244 | 
| [vanng822/go-premailer](https://github.com/vanng822/go-premailer) | Inline styling for html mail in golang | 82 | 
| [xhit/go-simple-mail](https://github.com/xhit/go-simple-mail) | Golang package for send email. Support keep alive connection, TLS and SSL. Easy for bulk SMTP. | 282 | 
| [hectane/hectane](https://github.com/hectane/hectane) | Lightweight SMTP client written in Go | 214 | 
| [matcornic/hermes](https://github.com/matcornic/hermes) | Golang package that generates clean, responsive HTML e-mails for sending transactional mail | 2444 | 
| [mailchain/mailchain](https://github.com/mailchain/mailchain) | Using Mailchain, blockchain users can now send and receive rich-media HTML messages with attachments via a blockchain address. | 115 | 
| [mailgun/mailgun-go](https://github.com/mailgun/mailgun-go) | Go library for sending mail with the Mailgun API. | 582 | 
| [mailhog/MailHog](https://github.com/mailhog/MailHog) | Web and API based SMTP testing | 9882 | 
| [valord577/mailx](https://github.com/valord577/mailx) | A library that makes it easier to send email via SMTP. | 2 | 
| [sendgrid/sendgrid-go](https://github.com/sendgrid/sendgrid-go) | The Official Twilio SendGrid Led, Community Driven Golang API Library | 801 | 
| [mailhog/smtp](https://github.com/mailhog/smtp) | MailHog SMTP Protocol | 70 | 
| [mocktools/go-smtp-mock](https://github.com/mocktools/go-smtp-mock) | SMTP mock server written on Golang. Mimic any 📤 SMTP server behaviour for your test environment with fake SMTP server. | 31 | 
| [truemail-rb/truemail-go](https://github.com/truemail-rb/truemail-go) | 🚀 Configurable Golang 📨 email validator/verifier. Verify email via Regex, DNS, SMTP and even more. | 5 | 


# Embeddable Scripting Languages

| Repository | Description | Stars |
|:---|:---|:---:|
| [mattn/anko](https://github.com/mattn/anko) | Scriptable interpreter written in golang | 1213 | 
| [alexeyco/binder](https://github.com/alexeyco/binder) | High level go to Lua binder. Write less, do more. | 56 | 
| [google/cel-go](https://github.com/google/cel-go) | Fast, portable, non-Turing complete expression evaluation with gradual typing (Go) | 1092 | 
| [krotik/ecal](https://github.com/krotik/ecal) | A simple embeddable scripting language which supports concurrent event processing. | 20 | 
| [antonmedv/expr](https://github.com/antonmedv/expr) | Expression language for Go | 2473 | 
| [gentee/gentee](https://github.com/gentee/gentee) | Gentee - script programming language for automation. It uses VM and compiler written in Go (Golang). | 89 | 
| [jcla1/gisp](https://github.com/jcla1/gisp) | Simple LISP in Go | 481 | 
| [olebedev/go-duktape](https://github.com/olebedev/go-duktape) | [abandoned] Duktape JavaScript engine bindings for Go | 779 | 
| [Shopify/go-lua](https://github.com/Shopify/go-lua) | A Lua VM in Go | 2294 | 
| [deuill/go-php](https://github.com/deuill/go-php) | PHP bindings for the Go programming language (Golang) | 853 | 
| [sbinet/go-python](https://github.com/sbinet/go-python) | naive go bindings to the CPython2 C-API | 1360 | 
| [dop251/goja](https://github.com/dop251/goja) | ECMAScript/JavaScript engine in pure Go | 2793 | 
| [aarzilli/golua](https://github.com/aarzilli/golua) | Go bindings for Lua C API - in progress | 583 | 
| [yuin/gopher-lua](https://github.com/yuin/gopher-lua) | GopherLua: VM and compiler for Lua in Go | 4618 | 
| [PaesslerAG/gval](https://github.com/PaesslerAG/gval) | Expression evaluation in golang | 467 | 
| [metacall/core](https://github.com/metacall/core) | MetaCall: The ultimate polyglot programming experience. | 910 | 
| [db47h/ngaro](https://github.com/db47h/ngaro) | An embeddable implementation of the Ngaro Virtual Machine for Go programs | 22 | 
| [ichiban/prolog](https://github.com/ichiban/prolog) | The only reasonable scripting engine for Go. | 357 | 
| [ian-kent/purl](https://github.com/ian-kent/purl) | Perl, but fluffy like a cat! | 34 | 
| [d5/tengo](https://github.com/d5/tengo) | A fast script language for Go | 2645 | 


# Error Handling

| Repository | Description | Stars |
|:---|:---|:---:|
| [emperror/emperror](https://github.com/emperror/emperror) | The Emperor takes care of all errors personally | 246 | 
| [rotisserie/eris](https://github.com/rotisserie/eris) | eris provides a better way to handle, trace, and log errors in Go 🎆 | 972 | 
| [snwfdhmp/errlog](https://github.com/snwfdhmp/errlog) | Reduce debugging time while programming Go. Use static and stack-trace analysis to determine which func call causes the error. | 402 | 
| [emperror/errors](https://github.com/emperror/errors) | Drop-in replacement for the standard library errors package and github.com/pkg/errors | 129 | 
| [pkg/errors](https://github.com/pkg/errors) | Simple error handling primitives | 7644 | 
| [neuronlabs/errors](https://github.com/neuronlabs/errors) | Simple golang error handling with classification primitives. | 3 | 
| [PumpkinSeed/errors](https://github.com/PumpkinSeed/errors) | Simple and efficient error package  | 4 | 
| [bnkamalesh/errors](https://github.com/bnkamalesh/errors) | A drop-in replacement for Go errors, with some added sugar! Unwrap user-friendly messages, HTTP status code, easy wrapping with multiple error types. | 32 | 
| [joomcode/errorx](https://github.com/joomcode/errorx) | A comprehensive error handling library for Go | 822 | 
| [ThundR67/falcon](https://github.com/ThundR67/falcon) | A Simple Yet Highly Powerful Package For Error Handling | 7 | 
| [hashicorp/go-multierror](https://github.com/hashicorp/go-multierror) | A Go (golang) package for representing a list of errors as a single error. | 1533 | 
| [ztrue/tracerr](https://github.com/ztrue/tracerr) | Golang errors with stack trace and source fragments. | 708 | 


# File Handling

| Repository | Description | Stars |
|:---|:---|:---:|
| [spf13/afero](https://github.com/spf13/afero) | A FileSystem Abstraction System for Go | 4326 | 
| [viant/afs](https://github.com/viant/afs) | Abstract File Storage | 174 | 
| [xis/baraka](https://github.com/xis/baraka) | a tool for handling file uploads simple | 41 | 
| [bigfile/bigfile](https://github.com/bigfile/bigfile) | Bigfile -- a file transfer system that supports http, rpc and ftp protocol   https://bigfile.site   | 219 | 
| [codingsince1985/checksum](https://github.com/codingsince1985/checksum) | Compute message digest for large files in Go | 56 | 
| [otiai10/copy](https://github.com/otiai10/copy) | Go copy directory recursively | 401 | 
| [homedepot/flop](https://github.com/homedepot/flop) | Go file operations library chasing GNU APIs. | 31 | 
| [dundee/gdu](https://github.com/dundee/gdu) | Fast disk usage analyzer with console interface written in Go | 1695 | 
| [artonge/go-csv-tag](https://github.com/artonge/go-csv-tag) | Read csv file from go using tags | 94 | 
| [hugocarreira/go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | copy files for humans | 16 | 
| [barasher/go-exiftool](https://github.com/barasher/go-exiftool) | Golang wrapper for Exiftool : extract as much metadata as possible (EXIF, ...) from files (pictures, pdf, office documents, ...) | 107 | 
| [artonge/go-gtfs](https://github.com/artonge/go-gtfs) | Load GTFS files in golang | 32 | 
| [no-src/gofs](https://github.com/no-src/gofs) | A file synchronization tool out of the box based on golang | 4 | 
| [1set/gut](https://github.com/1set/gut) | 🍱 yet another collection of go utilities & tools | 24 | 
| [dastoori/higgs](https://github.com/dastoori/higgs) | A tiny cross-platform Go library to hide/unhide files and directories | 9 | 
| [rjeczalik/notify](https://github.com/rjeczalik/notify) | File system event notification library on steroids. | 729 | 
| [qmuntal/opc](https://github.com/qmuntal/opc) | Go implementation of the Open Packaging Conventions (OPC) | 70 | 
| [parsyl/parquet](https://github.com/parsyl/parquet) | A library for reading and writing parquet files. | 50 | 
| [jonchun/pathtype](https://github.com/jonchun/pathtype) | Add a type for paths in Go. | 9 | 
| [pdfcpu/pdfcpu](https://github.com/pdfcpu/pdfcpu) | A PDF processor written in Go. | 3045 | 
| [dixonwille/skywalker](https://github.com/dixonwille/skywalker) | A package to allow one to concurrently go through a filesystem with ease | 71 | 
| [stl - Modules to read and write STL (stereolithography) files. Concurrent algorithm for reading.](https://gitlab.com/russoj88/stl?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [posener/tarfs](https://github.com/posener/tarfs) | An implementation of the FileSystem interface for tar files. | 49 | 
| [1set/todotxt](https://github.com/1set/todotxt) | Parser for todo.txt files in Go ✅ | 12 | 
| [C2FO/vfs](https://github.com/C2FO/vfs) | Pluggable, extensible virtual file system for Go | 161 | 


# Financial

| Repository | Description | Stars |
|:---|:---|:---:|
| [leekchan/accounting](https://github.com/leekchan/accounting) | money and currency formatting for golang | 724 | 
| [moov-io/ach](https://github.com/moov-io/ach) | ACH implements a reader, writer, and validator for Automated Clearing House (ACH) files. The HTTP server is available in a Docker image and the Go package is available. | 290 | 
| [bojanz/currency](https://github.com/bojanz/currency) | Currency handling for Go. | 281 | 
| [bnkamalesh/currency](https://github.com/bnkamalesh/currency) | A currency computations package. | 46 | 
| [shopspring/decimal](https://github.com/shopspring/decimal) | Arbitrary-precision fixed-point decimal numbers in go | 3838 | 
| [newity/fastme](https://github.com/newity/fastme) |  | 31 | 
| [FlashBoys/go-finance](https://github.com/FlashBoys/go-finance) | :warning: Deprecrated in favor of https://github.com/piquette/finance-go  | 536 | 
| [alpeb/go-finance](https://github.com/alpeb/go-finance) | Go library containing a collection of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. | 130 | 
| [pieterclaerhout/go-finance](https://github.com/pieterclaerhout/go-finance) | Finance related Go functions (e.g. exchange rates, VAT number checking, …) | 6 | 
| [m1/go-finnhub](https://github.com/m1/go-finnhub) | Simple and easy to use client for stock market, forex and crypto data from finnhub.io written in Go. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges | 67 | 
| [Rhymond/go-money](https://github.com/Rhymond/go-money) | Go implementation of Fowler's Money pattern | 1062 | 
| [aclindsa/ofxgo](https://github.com/aclindsa/ofxgo) | Golang library for querying and parsing OFX | 98 | 
| [i25959341/orderbook](https://github.com/i25959341/orderbook) | Matching Engine for Limit Order Book in Golang | 252 | 
| [jovandeginste/payme](https://github.com/jovandeginste/payme) | QR code generator (ASCII & PNG) for SEPA payments | 10 | 
| [BoltApp/sleet](https://github.com/BoltApp/sleet) | Payment abstraction library - one interface for multiple payment processors ( inspired by Ruby's ActiveMerchant ) | 85 | 
| [sdcoffey/techan](https://github.com/sdcoffey/techan) | Technical Analysis Library for Golang | 609 | 
| [achannarasappa/ticker](https://github.com/achannarasappa/ticker) | Terminal stock ticker with live updates and position tracking | 4179 | 
| [claygod/transaction](https://github.com/claygod/transaction) | Embedded database for accounts transactions. | 104 | 
| [dannyvankooten/vat](https://github.com/dannyvankooten/vat) | Go package for dealing with EU VAT. Does VAT number validation & rates retrieval. | 86 | 


# Forms

| Repository | Description | Stars |
|:---|:---|:---:|
| [robfig/bind](https://github.com/robfig/bind) |  | 27 | 
| [mholt/binding](https://github.com/mholt/binding) | Reflectionless data binding for Go's net/http (not actively maintained) | 789 | 
| [leebenson/conform](https://github.com/leebenson/conform) | Trims, sanitizes & scrubs data based on struct tags (go, golang) | 251 | 
| [go-playground/form](https://github.com/go-playground/form) | :steam_locomotive: Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. | 512 | 
| [monoculum/formam](https://github.com/monoculum/formam) | a package for decode form's values into struct in Go | 167 | 
| [albrow/forms](https://github.com/albrow/forms) | A lightweight go library for parsing form data or json from an http.Request. | 129 | 
| [gorilla/csrf](https://github.com/gorilla/csrf) | gorilla/csrf provides Cross Site Request Forgery (CSRF) prevention middleware for Go web applications & services 🔒 | 774 | 
| [ggicci/httpin](https://github.com/ggicci/httpin) | 🍡 HTTP Input for Go - Decode an HTTP request into a custom struct | 62 | 
| [justinas/nosurf](https://github.com/justinas/nosurf) | CSRF protection middleware for Go. | 1262 | 
| [sonh/qs](https://github.com/sonh/qs) | Go module for encoding structs into URL query parameters | 59 | 
| [TomWright/queryparam](https://github.com/TomWright/queryparam) | Go package to easily convert a URL's query parameters/values into usable struct values of the correct types. | 12 | 


# Functional

| Repository | Description | Stars |
|:---|:---|:---:|
| [TeaEntityLab/fpGo](https://github.com/TeaEntityLab/fpGo) | Monad, Functional Programming features for Golang | 245 | 
| [seborama/fuego](https://github.com/seborama/fuego) | Functional Experiment in Golang | 107 | 
| [tobyhede/go-underscore](https://github.com/tobyhede/go-underscore) |  Helpfully Functional Go -  A useful collection of Go utilities. Designed for programmer happiness.  | 1234 | 
| [rbrahul/gofp](https://github.com/rbrahul/gofp) | A super simple Lodash like utility library with essential functions that empowers the development in Go | 118 | 
| [rjNemo/underscore](https://github.com/rjNemo/underscore) | Useful functional programming helpers for Go 1.18 and beyond | 10 | 


# Game Development

| Repository | Description | Stars |
|:---|:---|:---:|
| [azul3d/engine](https://github.com/azul3d/engine) | Azul3D - A 3D game engine written in Go! | 532 | 
| [hajimehoshi/ebiten](https://github.com/hajimehoshi/ebiten) | A dead simple 2D game library for Go | 6194 | 
| [EngoEngine/engo](https://github.com/EngoEngine/engo) | Engo is an open-source 2D game engine written in Go. | 1486 | 
| [g3n/engine](https://github.com/g3n/engine) | Go 3D Game Engine (http://g3n.rocks) | 1878 | 
| [beefsack/go-astar](https://github.com/beefsack/go-astar) | Go implementation of the A* search algorithm | 488 | 
| [veandco/go-sdl2](https://github.com/veandco/go-sdl2) | SDL2 binding for Go | 1730 | 
| [ungerik/go3d](https://github.com/ungerik/go3d) | A performance oriented 2D/3D math package for Go | 254 | 
| [xtaci/gonet](https://github.com/xtaci/gonet) | A Game Server Skeleton in golang. | 1169 | 
| [xiaonanln/goworld](https://github.com/xiaonanln/goworld) | Scalable Distributed Game Server Engine with Hot Swapping in Golang | 2062 | 
| [name5566/leaf](https://github.com/name5566/leaf) | A game server framework in Go (golang) | 4330 | 
| [lonng/nano](https://github.com/lonng/nano) | Lightweight, facility, high performance golang based game server framework | 1968 | 
| [oakmound/oak](https://github.com/oakmound/oak) | A pure Go game engine | 1142 | 
| [topfreegames/pitaya](https://github.com/topfreegames/pitaya) | Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. | 1364 | 
| [faiface/pixel](https://github.com/faiface/pixel) | A hand-crafted 2D game library in Go | 3865 | 
| [gonutz/prototype](https://github.com/gonutz/prototype) | Simple 2D game prototyping framework. | 70 | 
| [gen2brain/raylib-go](https://github.com/gen2brain/raylib-go) | Go bindings for raylib, a simple and easy-to-use library to enjoy videogames programming. | 751 | 
| [JoelOtter/termloop](https://github.com/JoelOtter/termloop) | Terminal-based game engine for Go, built on top of Termbox | 1269 | 
| [kelindar/tile](https://github.com/kelindar/tile) | Tile is a 2D grid engine, built with data and cache friendly ways, includes pathfinding and observers. | 44 | 


# Generators

| Repository | Description | Stars |
|:---|:---|:---:|
| [t0pep0/efaceconv](https://github.com/t0pep0/efaceconv) |  | 51 | 
| [clipperhouse/gen](https://github.com/clipperhouse/gen) | Type-driven code generation for Go | 1358 | 
| [senselogic/GENERIS](https://github.com/senselogic/GENERIS) | Versatile Go code generator. | 31 | 
| [abice/go-enum](https://github.com/abice/go-enum) | An enum generator for go | 306 | 
| [ahmetb/go-linq](https://github.com/ahmetb/go-linq) | .NET LINQ capabilities in Go | 2896 | 
| [pieterclaerhout/go-xray](https://github.com/pieterclaerhout/go-xray) | Helpers for making the use of reflection easier | 21 | 
| [awalterschulze/goderive](https://github.com/awalterschulze/goderive) | Derives and generates mundane golang functions that you do not want to maintain yourself | 974 | 
| [wzshiming/gotype](https://github.com/wzshiming/gotype) | Golang source code parsing, usage like reflect package | 37 | 
| [jmattheis/goverter](https://github.com/jmattheis/goverter) | Generate type-safe Go converters by simply defining an interface | 105 | 
| [hexdigest/gowrap](https://github.com/hexdigest/gowrap) | GoWrap is a command line tool for generating decorators for Go interfaces | 592 | 
| [rjeczalik/interfaces](https://github.com/rjeczalik/interfaces) | Code generation tools for Go. | 334 | 
| [dave/jennifer](https://github.com/dave/jennifer) | Jennifer is a code generator for Go | 2355 | 
| [ungerik/pkgreflect](https://github.com/ungerik/pkgreflect) | A Go preprocessor for package scoped reflection | 99 | 
| [xiaoxin01/typeregistry](https://github.com/xiaoxin01/typeregistry) | create type dynamically in Golang | 13 | 


# Geographic

| Repository | Description | Stars |
|:---|:---|:---:|
| [hishamkaram/geoserver](https://github.com/hishamkaram/geoserver) | geoserver is a Go library for manipulating a GeoServer instance via the GeoServer REST API. | 70 | 
| [hishamkaram/gismanager](https://github.com/hishamkaram/gismanager) | Publish Your GIS Data(Vector Data) to PostGIS and Geoserver | 43 | 
| [airbusgeo/godal](https://github.com/airbusgeo/godal) | golang wrapper for github.com/OSGEO/gdal | 66 | 
| [consbio/mbtileserver](https://github.com/consbio/mbtileserver) | Basic Go server for mbtiles | 319 | 
| [paulmach/osm](https://github.com/paulmach/osm) | General purpose library for reading, writing and working with OpenStreetMap data | 201 | 
| [maguro/pbf](https://github.com/maguro/pbf) | OpenStreetMap PBF golang parser | 33 | 
| [pantrif/s2-geojson](https://github.com/pantrif/s2-geojson) | Draw a polygon on the map or paste a geoJSON and explore how the s2.RegionCoverer covers it with S2 cells depending on the min and max levels | 16 | 
| [golang/geo](https://github.com/golang/geo) | S2 geometry library in Go | 1335 | 
| [peterstace/simplefeatures](https://github.com/peterstace/simplefeatures) | Simple Features is a pure Go Implementation of the OpenGIS Simple Feature Access Specification | 48 | 
| [tidwall/tile38](https://github.com/tidwall/tile38) | Real-time Geospatial and Geofencing | 8018 | 
| [wroge/wgs84](https://github.com/wroge/wgs84) | A pure Go package for coordinate transformations. | 74 | 


# Go Compilers

| Repository | Description | Stars |
|:---|:---|:---:|
| [Konstantin8105/c4go](https://github.com/Konstantin8105/c4go) | Transpiling C code to Go code | 303 | 
| [andygeiss/esp32-transpiler](https://github.com/andygeiss/esp32-transpiler) | Transpile Golang into Arduino code to use fully automated testing at your IoT projects. | 42 | 
| [Konstantin8105/f4go](https://github.com/Konstantin8105/f4go) | Transpiling fortran code to golang code | 31 | 
| [gopherjs/gopherjs](https://github.com/gopherjs/gopherjs) | A compiler from Go to JavaScript for running Go code in a browser | 10967 | 
| [tardisgo/tardisgo](https://github.com/tardisgo/tardisgo) | Golang->Haxe->CPP/CSharp/Java/JavaScript transpiler   | 419 | 


# Goroutines

| Repository | Description | Stars |
|:---|:---|:---:|
| [panjf2000/ants](https://github.com/panjf2000/ants) | 🐜🐜🐜 ants is a high-performance and low-cost goroutine pool in Go, inspired by fasthttp./ ants 是一个高性能且低损耗的 goroutine 池。 | 7854 | 
| [mborders/artifex](https://github.com/mborders/artifex) | Simple in-memory job queue for Golang using worker-based dispatching | 129 | 
| [reugn/async](https://github.com/reugn/async) | Synchronization and asynchronous computation utilities library for Go | 40 | 
| [StudioSol/async](https://github.com/StudioSol/async) | A safe way to execute functions asynchronously, recovering them in case of panic. It also provides an error stack aiming to facilitate fail causes discovery. | 105 | 
| [kamilsk/breaker](https://github.com/kamilsk/breaker) | 🚧 Flexible mechanism to make execution flow interruptible. | 4 | 
| [ddelizia/channelify](https://github.com/ddelizia/channelify) | Make functions return a channel for parallel processing via go routines. | 20 | 
| [vivek-ng/concurrency-limiter](https://github.com/vivek-ng/concurrency-limiter) |  | 7 | 
| [ITcathyh/conexec](https://github.com/ITcathyh/conexec) | A concurrent toolkit to help execute funcs concurrently in an efficient and safe way. It supports specifying the overall timeout to avoid blocking. | 12 | 
| [marusama/cyclicbarrier](https://github.com/marusama/cyclicbarrier) | CyclicBarrier golang implementation | 93 | 
| [hexdigest/execpool](https://github.com/hexdigest/execpool) | A pool that spins up a given number of processes in advance and attaches stdin and stdout when needed. Very similar to FastCGI but works for any command. | 10 | 
| [workanator/go-floc](https://github.com/workanator/go-floc) | Floc: Orchestrate goroutines with ease. | 232 | 
| [kamildrazkiewicz/go-flow](https://github.com/kamildrazkiewicz/go-flow) | Simply way to control goroutines execution order based on dependencies | 186 | 
| [nikhilsaraf/go-tools](https://github.com/nikhilsaraf/go-tools) | A collection of tools for Golang | 7 | 
| [subchen/go-trylock](https://github.com/subchen/go-trylock) | TryLock support on read-write lock for Golang | 28 | 
| [pieterclaerhout/go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) | A sync.WaitGroup with error handling and concurrency control | 29 | 
| [catmullet/go-workers](https://github.com/catmullet/go-workers) | 👷 Library for safely running groups of workers concurrently or consecutively that require input and output through channels | 136 | 
| [zenthangplus/goccm](https://github.com/zenthangplus/goccm) | Limits the number of goroutines that are allowed to run concurrently | 47 | 
| [loveleshsharma/gohive](https://github.com/loveleshsharma/gohive) | 🐝 A Highly Performant and easy to use goroutine pool for Go | 32 | 
| [vardius/gollback](https://github.com/vardius/gollback) | Go asynchronous simple function utilities, for managing execution of closures and callbacks | 85 | 
| [hamed-yousefi/gowl](https://github.com/hamed-yousefi/gowl) | Gowl is a process management and process monitoring tool at once. An infinite worker pool gives you the ability to control the pool and processes and monitor their status. | 16 | 
| [benmanns/goworker](https://github.com/benmanns/goworker) | goworker is a Go-based background worker that runs 10 to 100,000* times faster than Ruby-based workers. | 2638 | 
| [xxjwxc/gowp](https://github.com/xxjwxc/gowp) | golang worker pool , Concurrency limiting goroutine pool | 364 | 
| [sherifabdlnaby/gpool](https://github.com/sherifabdlnaby/gpool) | gpool - a generic context-aware resizable goroutines pool to bound concurrency based on semaphore.  | 84 | 
| [ivpusic/grpool](https://github.com/ivpusic/grpool) | Lightweight Goroutine pool | 679 | 
| [duanckham/hands](https://github.com/duanckham/hands) | Hands is a process controller used to control the execution and return strategies of multiple goroutines. | 8 | 
| [AaronJan/Hunch](https://github.com/AaronJan/Hunch) | Hunch provides functions like: All, First, Retry, Waterfall etc., that makes asynchronous flow control more intuitive. | 77 | 
| [dirkaholic/kyoo](https://github.com/dirkaholic/kyoo) | Unlimited job queue for go, using a pool of concurrent workers processing the job queue entries | 37 | 
| [neilotoole/errgroup](https://github.com/neilotoole/errgroup) | errgroup with goroutine worker limits | 116 | 
| [arunsworld/nursery](https://github.com/arunsworld/nursery) | Structured Concurrency in Go | 40 | 
| [oversight - Oversight is a complete implementation of the Erlang supervision trees.](https://cirello.io/oversight?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [rafaeljesus/parallel-fn](https://github.com/rafaeljesus/parallel-fn) | Run functions in parallel :comet: | 33 | 
| [alitto/pond](https://github.com/alitto/pond) | 🔘 Minimalistic and High-performance goroutine worker pool written in Go | 477 | 
| [go-playground/pool](https://github.com/go-playground/pool) | :speedboat: a limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation | 673 | 
| [AnikHasibul/queue](https://github.com/AnikHasibul/queue) | package queue gives you a queue group accessibility. Helps you to limit goroutines, wait for the end of the all goroutines and much more. | 12 | 
| [x-mod/routine](https://github.com/x-mod/routine) | go routine control, abstraction of the Main and some useful Executors.如果你不会管理Goroutine的话，用它 | 48 | 
| [kamilsk/semaphore](https://github.com/kamilsk/semaphore) | 🚦 Semaphore pattern implementation with timeout of lock/unlock operations. | 90 | 
| [marusama/semaphore](https://github.com/marusama/semaphore) | Fast resizable golang semaphore primitive | 140 | 
| [ssgreg/stl](https://github.com/ssgreg/stl) | Software Transactional Locks | 24 | 
| [shettyh/threadpool](https://github.com/shettyh/threadpool) | Golang simple thread pool implementation | 69 | 
| [Jeffail/tunny](https://github.com/Jeffail/tunny) | A goroutine pool for Go | 2921 | 
| [vardius/worker-pool](https://github.com/vardius/worker-pool) | Go simple async worker pool | 82 | 
| [gammazero/workerpool](https://github.com/gammazero/workerpool) | Concurrency limiting goroutine pool | 751 | 


# GUI

| Repository | Description | Stars |
|:---|:---|:---:|


# Hardware

| Repository | Description | Stars |
|:---|:---|:---:|
| [arduino/arduino-cli](https://github.com/arduino/arduino-cli) | Arduino command line tool | 3331 | 
| [ziutek/emgo](https://github.com/ziutek/emgo) | Emgo: Bare metal Go (language for programming embedded systems) | 966 | 
| [jaypipes/ghw](https://github.com/jaypipes/ghw) | Golang hardware discovery/inspection library | 1127 | 
| [hypebeast/go-osc](https://github.com/hypebeast/go-osc) | Open Sound Control (OSC) library for Golang. Implemented in pure Go. | 144 | 
| [stianeikeland/go-rpio](https://github.com/stianeikeland/go-rpio) | :electric_plug: Raspberry Pi GPIO library for go-lang | 1837 | 
| [aler9/goroslib](https://github.com/aler9/goroslib) | ROS client library for the Go programming language | 172 | 
| [0xcafed00d/joystick](https://github.com/0xcafed00d/joystick) | Go Joystick API | 31 | 
| [zcalusic/sysinfo](https://github.com/zcalusic/sysinfo) | Sysinfo is a Go library providing Linux OS / kernel / hardware system information. | 328 | 


# Images

| Repository | Description | Stars |
|:---|:---|:---:|
| [anthonynsimon/bild](https://github.com/anthonynsimon/bild) | Image processing algorithms in pure Go | 3497 | 
| [h2non/bimg](https://github.com/h2non/bimg) | Go package for fast high-level image processing powered by libvips C library | 1840 | 
| [aofei/cameron](https://github.com/aofei/cameron) | An avatar generator for Go. | 82 | 
| [tdewolff/canvas](https://github.com/tdewolff/canvas) | Cairo in Go: vector to raster, SVG, PDF, EPS, WASM, OpenGL, Gio, etc. | 939 | 
| [gojek/darkroom](https://github.com/gojek/darkroom) |  | 186 | 
| [lucasepe/draft](https://github.com/lucasepe/draft) | Generate High Level Cloud Architecture diagrams using YAML syntax. | 522 | 
| [pravj/geopattern](https://github.com/pravj/geopattern) | :triangular_ruler: Create beautiful generative image patterns from a string in golang. | 1170 | 
| [fogleman/gg](https://github.com/fogleman/gg) | Go Graphics - 2D rendering in Go with a simple API. | 3322 | 
| [disintegration/gift](https://github.com/disintegration/gift) | Go Image Filtering Toolkit | 1541 | 
| [qmuntal/gltf](https://github.com/qmuntal/gltf) | :eyeglasses: Go library for encoding glTF 2.0 files | 150 | 
| [ungerik/go-cairo](https://github.com/ungerik/go-cairo) | Go binding for the cairo graphics library | 119 | 
| [bolknote/go-gd](https://github.com/bolknote/go-gd) | Go bingings for GD (http://www.boutell.com/gd/) | 53 | 
| [koyachi/go-nude](https://github.com/koyachi/go-nude) | Nudity detection with Go. | 345 | 
| [go-opencv/go-opencv](https://github.com/go-opencv/go-opencv) | Go bindings for OpenCV / 2.x API in gocv / 1.x API in opencv | 1269 | 
| [jyotiska/go-webcolors](https://github.com/jyotiska/go-webcolors) | Port of webcolors library from Python to Go | 26 | 
| [kolesa-team/go-webp](https://github.com/kolesa-team/go-webp) | Simple and fast webp library for golang | 46 | 
| [hybridgroup/gocv](https://github.com/hybridgroup/gocv) | Go package for computer vision using OpenCV 4 and beyond. | 4686 | 
| [corona10/goimagehash](https://github.com/corona10/goimagehash) | Go Perceptual image hashing package | 476 | 
| [corona10/goimghdr](https://github.com/corona10/goimghdr) | The imghdr module determines the type of image contained in a file for go | 38 | 
| [o1egl/govatar](https://github.com/o1egl/govatar) | Avatar generation library for GO language | 486 | 
| [davidbyttow/govips](https://github.com/davidbyttow/govips) | A lightning fast image processing and resizing library for Go | 651 | 
| [sensepost/gowitness](https://github.com/sensepost/gowitness) | 🔍 gowitness - a golang, web screenshot utility using Chrome Headless | 1401 | 
| [shomali11/gridder](https://github.com/shomali11/gridder) | A Grid based 2D Graphics library | 52 | 
| [qeesung/image2ascii](https://github.com/qeesung/image2ascii) | :foggy: Convert image to ASCII | 641 | 
| [gographics/imagick](https://github.com/gographics/imagick) | Go binding to ImageMagick's MagickWand C API | 1422 | 
| [h2non/imaginary](https://github.com/h2non/imaginary) | Fast, simple, scalable, Docker-ready HTTP microservice for high-level image processing | 4268 | 
| [disintegration/imaging](https://github.com/disintegration/imaging) | Imaging is a simple image processing package for Go | 4155 | 
| [hawx/img](https://github.com/hawx/img) | A selection of image manipulation tools | 138 | 
| [fogleman/ln](https://github.com/fogleman/ln) | 3D line art engine. | 3042 | 
| [noelyahan/mergi](https://github.com/noelyahan/mergi) | go library for image programming (merge, crop, resize, watermark, animate, ease, transit) | 166 | 
| [aldor007/mort](https://github.com/aldor007/mort) | Storage and image processing server written in Go | 448 | 
| [donatj/mpo](https://github.com/donatj/mpo) | JPEG-MPO Decoder / Converter Library and CLI Tool | 8 | 
| [thoas/picfit](https://github.com/thoas/picfit) | An image resizing server written in Go | 1624 | 
| [fogleman/pt](https://github.com/fogleman/pt) | A path tracer written in Go. | 2003 | 
| [nfnt/resize](https://github.com/nfnt/resize) | Pure golang image resizing  | 2797 | 
| [bamiaux/rez](https://github.com/bamiaux/rez) | Image resizing in pure Go and SIMD | 204 | 
| [jonoton/scout](https://github.com/jonoton/scout) | Scout is a standalone open source software solution for DIY video security. | 3 | 
| [muesli/smartcrop](https://github.com/muesli/smartcrop) | smartcrop finds good image crops for arbitrary crop sizes | 1613 | 
| [auyer/steganography](https://github.com/auyer/steganography) | Pure Golang Library that allows simple LSB steganography on images | 133 | 
| [DimitarPetrov/stegify](https://github.com/DimitarPetrov/stegify) | 🔍 Go tool for LSB steganography, capable of hiding any file within an image. | 996 | 
| [ajstarks/svgo](https://github.com/ajstarks/svgo) | Go Language Library for SVG generation | 1817 | 
| [ftrvxmtrx/tga](https://github.com/ftrvxmtrx/tga) | Go package for decoding and encoding TARGA image format | 30 | 
| [mehdipourfar/webp-server](https://github.com/mehdipourfar/webp-server) | Simple and minimal image server capable of storing, resizing, converting and caching images. | 38 | 


# IoT (Internet of Things)

| Repository | Description | Stars |
|:---|:---|:---:|
| [heedy/heedy](https://github.com/heedy/heedy) | An aggregator for personal metrics, and an extensible analysis engine | 323 | 
| [goiot/devices](https://github.com/goiot/devices) | Suite of libraries for IoT devices (written in Go), experimental for x/exp/io | 252 | 
| [xcodersun/eywa](https://github.com/xcodersun/eywa) | Make IoT a lot more fun with data.  | 52 | 
| [TIBCOSoftware/flogo](https://github.com/TIBCOSoftware/flogo) | Project Flogo is an open source ecosystem of opinionated  event-driven capabilities to simplify building efficient & modern serverless functions, microservices & edge apps. | 1944 | 
| [paypal/gatt](https://github.com/paypal/gatt) | Gatt is a Go package for building Bluetooth Low Energy peripherals | 1012 | 
| [hybridgroup/gobot](https://github.com/hybridgroup/gobot) | Golang framework for robotics, drones, and the Internet of Things (IoT) | 7659 | 
| [amimof/huego](https://github.com/amimof/huego) | An extensive Philips Hue client library for Go with an emphasis on simplicity | 201 | 
| [vaelen/iot](https://github.com/vaelen/iot) | A Go client for Google IoT Core | 56 | 
| [mainflux/mainflux](https://github.com/mainflux/mainflux) | Industrial IoT Messaging and Device Management Platform | 1727 | 
| [periph - Peripherals I/O to interface with low-level board facilities.](https://periph.io/?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [sensorbee/sensorbee](https://github.com/sensorbee/sensorbee) | Lightweight stream processing engine for IoT | 211 | 


# Job Scheduler

| Repository | Description | Stars |
|:---|:---|:---:|
| [deepaksinghvi/cdule](https://github.com/deepaksinghvi/cdule) | cdule (pronounce as Schedule) Golang based scheduler library with database support. | 2 | 
| [datarootsio/cheek](https://github.com/datarootsio/cheek) | Crontab-like scHeduler for Effective Execution of tasKs, cheek for short. | 25 | 
| [onatm/clockwerk](https://github.com/onatm/clockwerk) | Job Scheduling Library | 117 | 
| [krayzpipes/cronticker](https://github.com/krayzpipes/cronticker) | Golang ticker that works with Cron scheduling. | 2 | 
| [rk/go-cron](https://github.com/rk/go-cron) | A simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. | 210 | 
| [reugn/go-quartz](https://github.com/reugn/go-quartz) | Minimalist and zero-dependency scheduling library for Go | 614 | 
| [go-co-op/gocron](https://github.com/go-co-op/gocron) | Easy and fluent Go cron scheduling. This is a fork from https://github.com/jasonlvhit/gocron | 1706 | 
| [fieldryand/goflow](https://github.com/fieldryand/goflow) | Web UI-based workflow orchestrator for rapid prototyping | 18 | 
| [roylee0704/gron](https://github.com/roylee0704/gron) | gron, Cron Jobs in Go. | 887 | 
| [adhocore/gronx](https://github.com/adhocore/gronx) | Lightweight, fast and dependency-free Cron expression parser (due checker), task scheduler and/or daemon for Golang (tested on v1.13 and above) and standalone usage | 183 | 
| [bamzi/jobrunner](https://github.com/bamzi/jobrunner) | Framework for performing work asynchronously, outside of the request flow | 911 | 
| [albrow/jobs](https://github.com/albrow/jobs) | A persistent and flexible background jobs library for go. | 488 | 
| [kilgaloon/leprechaun](https://github.com/kilgaloon/leprechaun) | You had one job, or more then one, which can be done in steps | 84 | 
| [romshark/sched](https://github.com/romshark/sched) | A job scheduler for Go with the ability to fast-forward time. | 22 | 
| [carlescere/scheduler](https://github.com/carlescere/scheduler) | Job scheduling made easy. | 379 | 
| [madflojo/tasks](https://github.com/madflojo/tasks) | Package tasks is an easy to use in-process scheduler for recurring tasks in Go | 79 | 


# JSON

| Repository | Description | Stars |
|:---|:---|:---:|
| [spyzhov/ajson](https://github.com/spyzhov/ajson) | Abstract JSON for Golang with JSONPath support  | 109 | 
| [simonnilsson/ask](https://github.com/simonnilsson/ask) | A Go package that provides a simple way of accessing nested properties in maps and slices. | 14 | 
| [cocoonspace/dynjson](https://github.com/cocoonspace/dynjson) | Client-customizable JSON formats for dynamic APIs | 11 | 
| [lucassscaravelli/ej](https://github.com/lucassscaravelli/ej) | Write and read JSON from different sources in one line | 7 | 
| [vtopc/epoch](https://github.com/vtopc/epoch) | Contains primitives for marshaling/unmarshaling Unix timestamp/epoch to/from built-in time.Time type in JSON | 9 | 
| [valyala/fastjson](https://github.com/valyala/fastjson) | Fast JSON parser and validator for Go. No custom structs, no code generation, no reflection | 1475 | 
| [skanehira/gjo](https://github.com/skanehira/gjo) | Small utility to create JSON objects | 109 | 
| [tidwall/gjson](https://github.com/tidwall/gjson) | Get JSON values quickly - JSON parser for Go | 9934 | 
| [ddymko/go-jsonerror](https://github.com/ddymko/go-jsonerror) | Small package which wraps error responses to follow jsonapi.org | 12 | 
| [nicklaw5/go-respond](https://github.com/nicklaw5/go-respond) | A Go package for handling common HTTP JSON responses. | 45 | 
| [elgs/gojq](https://github.com/elgs/gojq) | JSON query in Golang | 183 | 
| [ChimeraCoder/gojson](https://github.com/ChimeraCoder/gojson) | Automatically generate Go (golang) struct definitions from example JSON | 2471 | 
| [yazgazan/jaydiff](https://github.com/yazgazan/jaydiff) | A JSON diff utility | 84 | 
| [wI2L/jettison](https://github.com/wI2L/jettison) | Highly configurable, fast JSON encoder for Go | 119 | 
| [romshark/jscan](https://github.com/romshark/jscan) | High performance JSON iterator for Go | 11 | 
| [JSON-to-Go - Convert JSON to Go struct.](https://mholt.github.io/json-to-go/?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [JSON-to-Proto - Convert JSON to Protobuf online.](https://json-to-proto.github.io/?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [m-zajac/json2go](https://github.com/m-zajac/json2go) | Create go type representation from json | 95 | 
| [AmuzaTkts/jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) | Go bindings based on the JSON API errors reference | 12 | 
| [neilotoole/jsoncolor](https://github.com/neilotoole/jsoncolor) | Colorized JSON output for Go | 28 | 
| [wI2L/jsondiff](https://github.com/wI2L/jsondiff) | Compute the diff between two JSON documents as a series of RFC6902 (JSON Patch) operations | 154 | 
| [miolini/jsonf](https://github.com/miolini/jsonf) | Console JSON formatter with query feature | 63 | 
| [ricardolonga/jsongo](https://github.com/ricardolonga/jsongo) | Fluent API to make it easier to create Json objects. | 100 | 
| [RichardKnop/jsonhal](https://github.com/RichardKnop/jsonhal) | A simple Go package to make custom structs marshal into HAL compatible JSON responses. | 10 | 
| [sinhashubham95/jsonic](https://github.com/sinhashubham95/jsonic) | All you need with JSON | 5 | 
| [zerosnake0/jzon](https://github.com/zerosnake0/jzon) | A golang json library inspired by jsoniter | 6 | 
| [qntfy/kazaam](https://github.com/qntfy/kazaam) | Arbitrary transformations of JSON in Golang | 222 | 
| [ake-persson/mapslice-json](https://github.com/ake-persson/mapslice-json) | Go MapSlice for ordered marshal/ unmarshal of maps in JSON | 9 | 
| [sanbornm/mp](https://github.com/sanbornm/mp) | Simple Email Parser | 44 | 
| [ohler55/ojg](https://github.com/ohler55/ojg) | Optimized JSON for Go | 459 | 
| [dedalqq/omg.jsonparser](https://github.com/dedalqq/omg.jsonparser) | The simple JSON parser with validation by condition | 3 | 
| [olvrng/ujson](https://github.com/olvrng/ujson) | µjson - A fast and minimal JSON parser and transformer that works on unstructured JSON | 54 | 
| [miladibra10/vjson](https://github.com/miladibra10/vjson) | vjson is a golang package that helps to validate JSON objects | 27 | 


# Logging

| Repository | Description | Stars |
|:---|:---|:---:|
| [amoghe/distillog](https://github.com/amoghe/distillog) | Logging, distilled | 26 | 
| [kpango/glg](https://github.com/kpango/glg) | Simple and blazing fast lockfree logging library for golang | 144 | 
| [lajosbencz/glo](https://github.com/lajosbencz/glo) | Logging library for Golang | 14 | 
| [golang/glog](https://github.com/golang/glog) | Leveled execution logs for Go | 3137 | 
| [utahta/go-cronowriter](https://github.com/utahta/go-cronowriter) | Time based rotating file writer | 47 | 
| [pieterclaerhout/go-log](https://github.com/pieterclaerhout/go-log) | A logging library with strack traces, object dumping and optional timestamps | 8 | 
| [subchen/go-log](https://github.com/subchen/go-log) | Simple and configurable Logging in Go, with level, formatters and writers | 11 | 
| [siddontang/go-log](https://github.com/siddontang/go-log) | a golang log lib supports level and multi handlers | 29 | 
| [ian-kent/go-log](https://github.com/ian-kent/go-log) | A logger, for Go | 37 | 
| [apsdehal/go-logger](https://github.com/apsdehal/go-logger) | Simple logger for Go programs. Allows custom formats for messages. | 274 | 
| [sadlil/gologger](https://github.com/sadlil/gologger) | The Simplest and worst logging library ever written | 40 | 
| [aphistic/gomol](https://github.com/aphistic/gomol) | Gomol is a library for structured, multiple-output logging for Go with extensible logging outputs | 18 | 
| [One-com/gone](https://github.com/One-com/gone) | Golang packages for writing small daemons and servers. | 40 | 
| [henvic/httpretty](https://github.com/henvic/httpretty) | Package httpretty prints the HTTP requests you make with Go pretty on your terminal. | 250 | 
| [ssgreg/journald](https://github.com/ssgreg/journald) | Go implementation of systemd Journal's native API for logging | 30 | 
| [clok/kemba](https://github.com/clok/kemba) | A tiny debug logging tool. Ideal for CLI tools and command applications. Inspired by https://github.com/visionmedia/debug | 6 | 
| [aerogo/log](https://github.com/aerogo/log) | :memo: Logging with multiple output targets. | 9 | 
| [apex/log](https://github.com/apex/log) | Structured logging package for Go. | 1217 | 
| [go-playground/log](https://github.com/go-playground/log) | :green_book: Simple, configurable and scalable Structured Logging for Go. | 277 | 
| [teris-io/log](https://github.com/teris-io/log) | Structured log interface | 24 | 
| [firstrow/logvoyage](https://github.com/firstrow/logvoyage) | LogVoyage - logging SaaS written in GoLang | 90 | 
| [inconshreveable/log15](https://github.com/inconshreveable/log15) | Structured, composable logging for Go | 1040 | 
| [ewwwwwqm/logdump](https://github.com/ewwwwwqm/logdump) | Package for multi-level logging | 9 | 
| [chzyer/logex](https://github.com/chzyer/logex) | An golang log lib, supports tracking and level, wrap by standard log lib | 38 | 
| [azer/logger](https://github.com/azer/logger) | Minimalistic logging library for Go. | 150 | 
| [mborders/logmatic](https://github.com/mborders/logmatic) | Colorized logger for Golang with dynamic log level configuration | 13 | 
| [mbndr/logo](https://github.com/mbndr/logo) | Golang logger to different configurable writers. | 9 | 
| [sirupsen/logrus](https://github.com/sirupsen/logrus) | Structured, pluggable logging for Go. | 20165 | 
| [cabify/logrusiowriter](https://github.com/cabify/logrusiowriter) | io.Writer implementation using logrus logger | 13 | 
| [sebest/logrusly](https://github.com/sebest/logrusly) | Loggly Hooks for GO Logrus logger | 27 | 
| [logur/logur](https://github.com/logur/logur) | Logur is an opinionated collection of logging best practices | 154 | 
| [hashicorp/logutils](https://github.com/hashicorp/logutils) | Utilities for slightly better logging in Go (Golang). | 317 | 
| [mgutz/logxi](https://github.com/mgutz/logxi) | A 12-factor app logger built for performance and happy development | 348 | 
| [natefinch/lumberjack](https://github.com/natefinch/lumberjack) | lumberjack is a log rolling package for Go | 3183 | 
| [jbrodriguez/mlog](https://github.com/jbrodriguez/mlog) | A simple logging module for go, with a rotating file feature and console logging. | 24 | 
| [gyozatech/noodlog](https://github.com/gyozatech/noodlog) | 🍜 Parametrized JSON logging library in Golang which lets you obfuscate sensitive data and marshal any kind of content. | 36 | 
| [francoispqt/onelog](https://github.com/francoispqt/onelog) | Dead simple, super fast, zero allocation and modular logger for Golang | 402 | 
| [go-ozzo/ozzo-log](https://github.com/go-ozzo/ozzo-log) | A Go (golang) package providing high-performance asynchronous logging, message filtering by severity and category, and multiple message targets. | 118 | 
| [phuslu/log](https://github.com/phuslu/log) | Structured Logging Made Easy | 419 | 
| [arthurkiller/rollingwriter](https://github.com/arthurkiller/rollingwriter) | Rolling writer is an IO util for auto rolling write in go. | 228 | 
| [cihub/seelog](https://github.com/cihub/seelog) | Seelog is a native Go logging library that provides flexible asynchronous dispatching, filtering, and formatting. | 1594 | 
| [davecgh/go-spew](https://github.com/davecgh/go-spew) | Implements a deep pretty printer for Go data structures to aid in debugging | 4924 | 
| [simukti/sqldb-logger](https://github.com/simukti/sqldb-logger) | A logger for Go SQL database driver without modifying existing *sql.DB stdlib usage. | 213 | 
| [alexcesaro/log](https://github.com/alexcesaro/log) | Logging packages for Go | 45 | 
| [structy/log](https://github.com/structy/log) | A simple to use log system, minimalist but with features for debugging and differentiation of messages | 4 | 
| [hpcloud/tail](https://github.com/hpcloud/tail) | Go package for reading from continously updated files (tail -f) | 2297 | 
| [xfxdev/xlog](https://github.com/xfxdev/xlog) | plugin architecture and flexible log system for golang | 6 | 
| [rs/xlog](https://github.com/rs/xlog) | xlog is a logger for net/context aware HTTP applications | 135 | 
| [jfcg/yell](https://github.com/jfcg/yell) | :ledger: Yet another minimalist logging library | 0 | 
| [uber-go/zap](https://github.com/uber-go/zap) | Blazing fast, structured, leveled logging in Go. | 15260 | 
| [rs/zerolog](https://github.com/rs/zerolog) | Zero Allocation JSON Logger | 6012 | 
| [edoger/zkits-logger](https://github.com/edoger/zkits-logger) | A powerful zero-dependency json logger. | 15 | 


# Machine Learning

| Repository | Description | Stars |
|:---|:---|:---:|
| [jbrukh/bayesian](https://github.com/jbrukh/bayesian) | Naive Bayesian Classification for Golang. | 729 | 
| [ryanbressler/CloudForest](https://github.com/ryanbressler/CloudForest) | Ensembles of decision trees in go/golang. | 708 | 
| [sgrodriguez/ddt](https://github.com/sgrodriguez/ddt) | Golang Dynamic Decision Tree | 16 | 
| [MaxHalford/eaopt](https://github.com/MaxHalford/eaopt) | :four_leaf_clover: Evolutionary optimization library for Go (genetic algorithm, partical swarm optimization, differential evolution) | 780 | 
| [khezen/evoli](https://github.com/khezen/evoli) | Genetic Algorithm and Particle Swarm Optimization | 22 | 
| [Fontinalis/fonet](https://github.com/Fontinalis/fonet) | fonet is a deep neural network package for Go. | 64 | 
| [e-XpertSolutions/go-cluster](https://github.com/e-XpertSolutions/go-cluster) | k-modes and k-prototypes clustering algorithms implementation in Go | 31 | 
| [patrikeh/go-deep](https://github.com/patrikeh/go-deep) | Artificial Neural Network | 357 | 
| [vksnk/go-fann](https://github.com/vksnk/go-fann) | Go bindings for FANN, library for artificial neural networks | 108 | 
| [nikolaydubina/go-featureprocessing](https://github.com/nikolaydubina/go-featureprocessing) | 🔥 Fast, simple sklearn-like feature processing for Go | 71 | 
| [thoj/go-galib](https://github.com/thoj/go-galib) | Genetic Algorithms library written in Go / golang | 189 | 
| [daviddengcn/go-pr](https://github.com/daviddengcn/go-pr) | Pattern recognition package in Go lang. | 61 | 
| [goml/gobrain](https://github.com/goml/gobrain) | Neural Networks written in go | 518 | 
| [e-dard/godist](https://github.com/e-dard/godist) | Probability distributions and associated methods in Go | 32 | 
| [tomcraven/goga](https://github.com/tomcraven/goga) | Golang Genetic Algorithm | 152 | 
| [sjwhitworth/golearn](https://github.com/sjwhitworth/golearn) | Machine Learning for Go | 8250 | 
| [danieldk/golinear](https://github.com/danieldk/golinear) | liblinear bindings for Go | 44 | 
| [surenderthakran/gomind](https://github.com/surenderthakran/gomind) | A simplistic Neural Network Library in Go | 23 | 
| [cdipaolo/goml](https://github.com/cdipaolo/goml) | On-line Machine Learning in Go (and so much more) | 1320 | 
| [dathoangnd/gonet](https://github.com/dathoangnd/gonet) | Neural Network for Go. | 73 | 
| [c-bata/goptuna](https://github.com/c-bata/goptuna) | A hyperparameter optimization framework, inspired by Optuna. | 208 | 
| [timkaye11/goRecommend](https://github.com/timkaye11/goRecommend) | Collaborative Filtering (CF) Algorithms in Go!  | 182 | 
| [gorgonia/gorgonia](https://github.com/gorgonia/gorgonia) | Gorgonia is a library that helps facilitate machine learning in Go. | 4434 | 
| [zhenghaoz/gorse](https://github.com/zhenghaoz/gorse) | An open source recommender system service written in Go | 5335 | 
| [asafschers/goscore](https://github.com/asafschers/goscore) | Go Scoring API for PMML | 75 | 
| [otiai10/gosseract](https://github.com/otiai10/gosseract) | Go package for OCR (Optical Character Recognition), by using Tesseract C++ library | 1700 | 
| [datastream/libsvm](https://github.com/datastream/libsvm) | libsvm go version | 71 | 
| [BayesWitnesses/m2cgen](https://github.com/BayesWitnesses/m2cgen) | Transform ML models into a native code (Java, C, Python, Go, JavaScript, Visual Basic, C#, R, PowerShell, PHP, Dart, Haskell, Ruby, F#, Rust) with zero dependencies | 2063 | 
| [jinyeom/neat](https://github.com/jinyeom/neat) | NEAT (NeuroEvolution of Augmenting Topologies) implemented in Go | 63 | 
| [schuyler/neural-go](https://github.com/schuyler/neural-go) | A multilayer perceptron network implemented in Go, with training via backpropagation. | 63 | 
| [otiai10/ocrserver](https://github.com/otiai10/ocrserver) | A simple OCR API server, seriously easy to be deployed by Docker, on Heroku as well | 479 | 
| [owulveryck/onnx-go](https://github.com/owulveryck/onnx-go) | onnx-go gives the ability to import a pre-trained neural network within Go without being linked to a framework or library. | 381 | 
| [ThePaw/probab](https://github.com/ThePaw/probab) | Automatically exported from code.google.com/p/probab | 17 | 
| [malaschitz/randomForest](https://github.com/malaschitz/randomForest) | Random Forest implementation in golang | 20 | 
| [muesli/regommend](https://github.com/muesli/regommend) | Recommendation engine for Go | 299 | 
| [eaigner/shield](https://github.com/eaigner/shield) | Bayesian text classifier with flexible tokenizers and storage backends for Go | 150 | 
| [galeone/tfgo](https://github.com/galeone/tfgo) | Tensorflow + Go, the gopher way | 1907 | 
| [Xamber/Varis](https://github.com/Xamber/Varis) | Golang Neural Network  | 44 | 


# Messaging

| Repository | Description | Stars |
|:---|:---|:---:|
| [kak-tus/ami](https://github.com/kak-tus/ami) | Go client to reliable queues based on Redis Cluster Streams | 23 | 
| [rabbitmq/amqp091-go](https://github.com/rabbitmq/amqp091-go) | An AMQP 0-9-1 Go client maintained by the RabbitMQ team. Originally by @streadway: `streadway/amqp` | 279 | 
| [sideshow/apns2](https://github.com/sideshow/apns2) | ⚡ HTTP/2 Apple Push Notification Service (APNs) push provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps, using the APNs HTTP/2 protocol. | 2604 | 
| [hibiken/asynq](https://github.com/hibiken/asynq) | Simple, reliable, and efficient distributed task queue in Go | 2872 | 
| [Clivern/Beaver](https://github.com/Clivern/Beaver) | 💨 A real time messaging system to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. | 1318 | 
| [benthosdev/benthos](https://github.com/benthosdev/benthos) | Fancy stream processing made operationally mundane | 4209 | 
| [mustafaturan/bus](https://github.com/mustafaturan/bus) | 🔊Minimalist message bus implementation for internal communication with zero-allocation magic on Emit | 250 | 
| [centrifugal/centrifugo](https://github.com/centrifugal/centrifugo) | Scalable real-time messaging server in a language-agnostic way. Set up once and forever. | 5903 | 
| [chanify/chanify](https://github.com/chanify/chanify) | Chanify is a safe and simple notification tools. This repository is command line tools for Chanify. | 787 | 
| [jeroenrinzema/commander](https://github.com/jeroenrinzema/commander) | Build event-driven and event streaming applications with ease | 59 | 
| [confluentinc/confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) | Confluent's Apache Kafka Golang client | 3195 | 
| [godbus/dbus](https://github.com/godbus/dbus) | Native Go bindings for D-Bus | 693 | 
| [appleboy/drone-line](https://github.com/appleboy/drone-line) | Sending line notifications using a binary, docker or Drone CI. | 76 | 
| [olebedev/emitter](https://github.com/olebedev/emitter) | Emits events in Go way, with wildcard, predicates, cancellation possibilities and many other good wins | 419 | 
| [agoalofalife/event](https://github.com/agoalofalife/event) | The implementation of the pattern observer | 45 | 
| [asaskevich/EventBus](https://github.com/asaskevich/EventBus) | [Go] Lightweight eventbus with async compatibility for Go | 1108 | 
| [osamingo/gaurun-client](https://github.com/osamingo/gaurun-client) | Gaurun Client written in Go | 10 | 
| [desertbit/glue](https://github.com/desertbit/glue) | Glue - Robust Go and Javascript Socket Library (Alternative to Socket.io) | 401 | 
| [cheshir/go-mq](https://github.com/cheshir/go-mq) | Declare AMQP entities like queues, producers, and consumers in a declarative way. Can be used to work with RabbitMQ. | 73 | 
| [TheCreeper/go-notify](https://github.com/TheCreeper/go-notify) | Package notify provides an implementation of the Gnome DBus Notifications Specification. | 59 | 
| [nsqio/go-nsq](https://github.com/nsqio/go-nsq) | The official Go package for NSQ | 2110 | 
| [jirenius/go-res](https://github.com/jirenius/go-res) | RES Service protocol library for Go | 54 | 
| [googollee/go-socket.io](https://github.com/googollee/go-socket.io) | socket.io library for golang, a realtime application framework. | 4522 | 
| [maxatome/go-vitotrol](https://github.com/maxatome/go-vitotrol) | golang client library to Viessmann Vitotrol web service | 17 | 
| [trivago/gollum](https://github.com/trivago/gollum) | An n:m message multiplexer written in Go | 917 | 
| [jcuga/golongpoll](https://github.com/jcuga/golongpoll) | golang long polling library.  Makes web pub-sub easy via HTTP long-poll servers and clients :smiley: :coffee: :computer: | 594 | 
| [Terry-Mao/gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) | Golang push server cluster | 2025 | 
| [appleboy/gorush](https://github.com/appleboy/gorush) | A push notification server written in Go (Golang). | 6169 | 
| [alexsniffin/gosd](https://github.com/alexsniffin/gosd) | A library for scheduling when to dispatch a message to a channel | 19 | 
| [smancke/guble](https://github.com/smancke/guble) | websocket based messaging server written in golang | 151 | 
| [leozz37/hare](https://github.com/leozz37/hare) | 🐇  CLI tool for websockets and easy to use Golang package | 38 | 
| [leandro-lugaresi/hub](https://github.com/leandro-lugaresi/hub) | :incoming_envelope: A fast Message/Event Hub using publish/subscribe pattern with support for topics like* rabbitMQ exchanges for Go applications | 114 | 
| [socifi/jazz](https://github.com/socifi/jazz) | Abstraction layer for simple rabbitMQ connection, messaging and administration | 15 | 
| [RichardKnop/machinery](https://github.com/RichardKnop/machinery) | Machinery is an asynchronous task queue/job queue based on distributed message passing. | 6065 | 
| [nanomsg/mangos](https://github.com/nanomsg/mangos) | mangos is a pure Golang implementation of nanomsg's "Scalablilty Protocols" | 489 | 
| [olahol/melody](https://github.com/olahol/melody) | :notes: Minimalist websocket framework for Go | 2390 | 
| [dunglas/mercure](https://github.com/dunglas/mercure) | Server-sent live updates: protocol and reference implementation | 2679 | 
| [vardius/message-bus](https://github.com/vardius/message-bus) | Go simple async message bus | 209 | 
| [nats-io/nats.go](https://github.com/nats-io/nats.go) | Golang client for NATS, the cloud native messaging system. | 3844 | 
| [rafaeljesus/nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) | A tiny wrapper around NSQ topic and channel :rocket: | 71 | 
| [dailymotion/oplog](https://github.com/dailymotion/oplog) | A generic oplog/replication system for microservices | 111 | 
| [cskr/pubsub](https://github.com/cskr/pubsub) | A simple pubsub package for go. | 370 | 
| [rafaeljesus/rabbus](https://github.com/rafaeljesus/rabbus) | A tiny wrapper over amqp exchanges and queues 🚌 ✨ | 94 | 
| [jandelgado/rabtap](https://github.com/jandelgado/rabtap) | RabbitMQ wire tap and swiss army knife | 209 | 
| [sybrexsys/RapidMQ](https://github.com/sybrexsys/RapidMQ) | RapidMQ is a pure, extremely productive, lightweight and reliable library for managing of the local messages queue | 64 | 
| [robinjoseph08/redisqueue](https://github.com/robinjoseph08/redisqueue) | redisqueue provides a producer and consumer of a queue that uses Redis streams | 80 | 
| [sbabiv/rmqconn](https://github.com/sbabiv/rmqconn) | RabbitMQ Reconnection client | 17 | 
| [Shopify/sarama](https://github.com/Shopify/sarama) | Sarama is a Go library for Apache Kafka. | 8328 | 
| [uniqush/uniqush-push](https://github.com/uniqush/uniqush-push) | Uniqush is a free and open source software system which provides a unified push service for server side notification to apps on mobile devices. | 1328 | 
| [pebbe/zmq4](https://github.com/pebbe/zmq4) | A Go interface to ZeroMQ version 4 | 971 | 


# Microsoft Office

| Repository | Description | Stars |
|:---|:---|:---:|


# Microsoft Excel

| Repository | Description | Stars |
|:---|:---|:---:|
| [qax-os/excelize](https://github.com/qax-os/excelize) | Go language library for reading and writing Microsoft Excel™ (XLSX) files. | 11125 | 
| [szyhf/go-excel](https://github.com/szyhf/go-excel) | A simple and light excel file reader to read a standard excel as a table faster | 一个轻量级的Excel数据读取库，用一种更`关系数据库`的方式解析Excel。 | 139 | 
| [fterrag/goxlsxwriter](https://github.com/fterrag/goxlsxwriter) | Golang bindings for libxlsxwriter for writing XLSX files | 18 | 
| [tealeg/xlsx](https://github.com/tealeg/xlsx) | Go (golang) library for reading and writing XLSX files.  | 5265 | 
| [plandem/xlsx](https://github.com/plandem/xlsx) | Fast and reliable way to work with Microsoft Excel™ [xlsx] files in Golang | 148 | 


# Miscellaneous

| Repository | Description | Stars |
|:---|:---|:---:|


# Dependency Injection

| Repository | Description | Stars |
|:---|:---|:---:|
| [magic003/alice](https://github.com/magic003/alice) | An additive dependency injection container for Golang. | 44 | 
| [golobby/container](https://github.com/golobby/container) | A lightweight yet powerful IoC dependency injection container for Go projects | 316 | 
| [goava/di](https://github.com/goava/di) | 🛠 A full-featured dependency injection container for go programming language. | 145 | 
| [uber-go/dig](https://github.com/uber-go/dig) | A reflection based dependency injection toolkit for Go. | 2425 | 
| [i-love-flamingo/dingo](https://github.com/i-love-flamingo/dingo) | Go Dependency Injection Framework | 124 | 
| [uber-go/fx](https://github.com/uber-go/fx) | A dependency injection based application framework for Go. | 2624 | 
| [vardius/gocontainer](https://github.com/vardius/gocontainer) | Simple Dependency Injection Container | 15 | 
| [goioc/di](https://github.com/goioc/di) | Simple and yet powerful Dependency Injection for Go | 153 | 
| [google/wire](https://github.com/google/wire) | Compile-time Dependency Injection for Go | 7766 | 
| [HnH/di](https://github.com/HnH/di) | DI container library that is focused on clean API and flexibility. | 3 | 
| [go-kata/kinit](https://github.com/go-kata/kinit) | GO Dependency Injection | 5 | 
| [logrange/linker](https://github.com/logrange/linker) | Dependency Injection and Inversion of Control package | 32 | 
| [muir/nject](https://github.com/muir/nject) | Golang type-safe dependency injection | 5 | 
| [Fs02/wire](https://github.com/Fs02/wire) | Strict Runtime Dependency Injection for Golang | 34 | 


# Project Layout

| Repository | Description | Stars |
|:---|:---|:---:|
| [ardanlabs/service](https://github.com/ardanlabs/service) | Starter code for writing web services in Go using Kubernetes. | 2253 | 
| [lacion/cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) | A Go project template | 508 | 
| [zitryss/go-sample](https://github.com/zitryss/go-sample) | Go Project Sample Layout | 96 | 
| [allaboutapps/go-starter](https://github.com/allaboutapps/go-starter) | An opinionated production-ready SQL-/Swagger-first RESTful JSON API written in Go, highly integrated with VSCode DevContainers by allaboutapps. | 139 | 
| [Fs02/go-todo-backend](https://github.com/Fs02/go-todo-backend) | Go Todo Backend example using modular project layout for product microservice. | 134 | 
| [wajox/gobase](https://github.com/wajox/gobase) | This is a simple skeleton for golang applications | 23 | 
| [golang-standards/project-layout](https://github.com/golang-standards/project-layout) | Standard Go Project Layout | 30436 | 
| [golang-templates/seed](https://github.com/golang-templates/seed) | Go application GitHub repository template. | 252 | 
| [insidieux/inizio](https://github.com/insidieux/inizio) | Golang project standard layout generator | 9 | 
| [sagikazarmark/modern-go-application](https://github.com/sagikazarmark/modern-go-application) | Modern Go Application example | 1160 | 
| [mikestefanello/pagoda](https://github.com/mikestefanello/pagoda) | Rapid, easy full-stack web development starter kit in Go | 171 | 
| [catchplay/scaffold](https://github.com/catchplay/scaffold) | Generate scaffold project layout for Go. | 108 | 
| [wangyoucao577/go-project-layout](https://github.com/wangyoucao577/go-project-layout) | My understanding of how to structure a golang project.  | 16 | 


# Strings

| Repository | Description | Stars |
|:---|:---|:---:|
| [mkungla/bexp](https://github.com/mkungla/bexp) | Go implementation of Brace Expansion mechanism to generate arbitrary strings. | 6 | 
| [go-formatter - Implements replacement fields surrounded by curly braces {} format strings.](https://gitlab.com/tymonx/go-formatter?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [gobeam/stringy](https://github.com/gobeam/stringy) | Convert string to camel case, snake case, kebab case / slugify, custom delimiter, pad string, tease string and many other functionalities with help of by Stringy package. | 109 | 
| [ozgio/strutil](https://github.com/ozgio/strutil) | String utilities for Go | 161 | 
| [abhimanyu003/sttr](https://github.com/abhimanyu003/sttr) | cross-platform, cli app to perform various operations on string | 350 | 
| [huandu/xstrings](https://github.com/huandu/xstrings) | Implements string functions widely used in other languages but absent in Go. | 1002 | 


# Uncategorized

| Repository | Description | Stars |
|:---|:---|:---:|
| [mudler/anagent](https://github.com/mudler/anagent) | Minimalistic, pluggable Golang evloop/timer handler with dependency-injection | 14 | 
| [antchfx/antch](https://github.com/antchfx/antch) | Antch, a fast, powerful and extensible web crawling & scraping framework for Go | 227 | 
| [mholt/archiver](https://github.com/mholt/archiver) | Easily create & extract archives, and compress & decompress files of various formats | 3512 | 
| [artyom/autoflags](https://github.com/artyom/autoflags) | Populate go command line app flags from config struct | 36 | 
| [kirillDanshin/avgRating](https://github.com/kirillDanshin/avgRating) | Calculate average score and rating based on Wilson Score Equation | 11 | 
| [dimiro1/banner](https://github.com/dimiro1/banner) | An easy way to add useful startup banners into your Go applications | 381 | 
| [mojocn/base64Captcha](https://github.com/mojocn/base64Captcha) | captcha of base64 image string | 1391 | 
| [bobg/basexx](https://github.com/bobg/basexx) | Convert digit strings between arbitrary bases. | 2 | 
| [distatus/battery](https://github.com/distatus/battery) | cross-platform, normalized battery information library | 201 | 
| [icza/bitio](https://github.com/icza/bitio) | Optimized bit-level Reader and Writer for Go. | 183 | 
| [digitalcrab/browscap_go](https://github.com/digitalcrab/browscap_go) | GoLang Library for Browser Capabilities Project | 40 | 
| [steambap/captcha](https://github.com/steambap/captcha) | :sunglasses:Package captcha provides an easy to use, unopinionated API for captcha generation | 96 | 
| [cstockton/go-conv](https://github.com/cstockton/go-conv) | Fast conversions across various Go types with a simple API. | 376 | 
| [miolini/datacounter](https://github.com/miolini/datacounter) | Golang counters for readers/writers | 37 | 
| [neotoolkit/faker](https://github.com/neotoolkit/faker) | Fake data generator | 3 | 
| [pioz/faker](https://github.com/pioz/faker) | Random fake data and struct generator for Go. | 60 | 
| [go-ffmt/ffmt](https://github.com/go-ffmt/ffmt) | Golang beautify data display for Humans | 258 | 
| [TwiN/gatus](https://github.com/TwiN/gatus) | ⛑ Gatus - Automated service health dashboard | 2365 | 
| [lana/go-commandbus](https://github.com/lana/go-commandbus) | Simple command bus for GO | 5 | 
| [jolestar/go-commons-pool](https://github.com/jolestar/go-commons-pool) | a generic object pool for golang | 1020 | 
| [go-openapi - Collection of packages to parse and utilize open-api schemas.](https://github.com/go-openapi?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [eapache/go-resiliency](https://github.com/eapache/go-resiliency) | Resiliency patterns for golang | 1341 | 
| [gen2brain/go-unarr](https://github.com/gen2brain/go-unarr) | Go bindings for unarr (decompression library for RAR, TAR, ZIP and 7z archives) | 174 | 
| [brianvoe/gofakeit](https://github.com/brianvoe/gofakeit) | Random fake data generator written in go | 2333 | 
| [antham/gommit](https://github.com/antham/gommit) | Enforce git message commit consistency | 95 | 
| [shirou/gopsutil](https://github.com/shirou/gopsutil) | psutil for golang | 7488 | 
| [osamingo/gosh](https://github.com/osamingo/gosh) | Provide Go Statistics Handler, Struct, Measure Method | 29 | 
| [haxpax/gosms](https://github.com/haxpax/gosms) | :mailbox_closed: Your own local SMS gateway in Go | 1373 | 
| [cabify/gotoprom](https://github.com/cabify/gotoprom) | Type-safe Prometheus metrics builder library for golang | 93 | 
| [pariz/gountries](https://github.com/pariz/gountries) | Gountries provides: Countries (ISO-3166-1), Country Subdivisions(ISO-3166-2), Currencies (ISO 4217), Geo Coordinates(ISO-6709) as well as translations, country borders and other stuff exposed as struct data. | 337 | 
| [ddddddO/gtree](https://github.com/ddddddO/gtree) | Output tree🌳 or Make directories(files)📁 from Markdown or Programmatically. Provide CLI and Go Package. | 35 | 
| [alexliesenfeld/health](https://github.com/alexliesenfeld/health) | A simple and flexible health check library for Go. | 489 | 
| [dimiro1/health](https://github.com/dimiro1/health) | An easy to use, extensible health check library for Go applications. | 429 | 
| [etherlabsio/healthcheck](https://github.com/etherlabsio/healthcheck) | An simple, easily extensible and concurrent health-check library for Go services | 214 | 
| [Wing924/hostutils](https://github.com/Wing924/hostutils) | A golang library for packing and unpacking hosts list | 10 | 
| [osamingo/indigo](https://github.com/osamingo/indigo) | A distributed unique ID generator of using Sonyflake and encoded by Base58 | 93 | 
| [hyperboloide/lk](https://github.com/hyperboloide/lk) | Simple licensing library for golang. | 236 | 
| [llir/llvm](https://github.com/llir/llvm) | Library for interacting with LLVM IR in pure Go. | 866 | 
| [pascaldekloe/metrics](https://github.com/pascaldekloe/metrics) | atomic measures + Prometheus exposition library | 22 | 
| [alwindoss/morse](https://github.com/alwindoss/morse) | Morse Code Library in Go | 75 | 
| [lrita/numa](https://github.com/lrita/numa) | NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code. | 8 | 
| [neotoolkit/openapi](https://github.com/neotoolkit/openapi) | OpenAPI 3.x parser | 5 | 
| [hyperboloide/pdfgen](https://github.com/hyperboloide/pdfgen) | HTTP service to generate PDF from Json requests | 56 | 
| [mavihq/persian](https://github.com/mavihq/persian) | Some utilities for Persian language in Go (Golang) | 62 | 
| [aofei/sandid](https://github.com/aofei/sandid) | Every grain of sand on Earth has its own ID. | 33 | 
| [Wing924/shellwords](https://github.com/Wing924/shellwords) | A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. | 17 | 
| [teris-io/shortid](https://github.com/teris-io/shortid) | Super short, fully unique, non-sequential and URL friendly Ids | 749 | 
| [containrrr/shoutrrr](https://github.com/containrrr/shoutrrr) | Notification library for gophers and their furry friends. | 336 | 
| [qmuntal/stateless](https://github.com/qmuntal/stateless) | Go library for creating state machines | 412 | 
| [go-playground/stats](https://github.com/go-playground/stats) | :chart_with_upwards_trend: Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... | 158 | 
| [hackebrot/turtle](https://github.com/hackebrot/turtle) | Emojis for Go 😄🐢🚀 | 131 | 
| [pantrif/url-shortener](https://github.com/pantrif/url-shortener) | A golang URL Shortener | 35 | 
| [azr/generators](https://github.com/azr/generators) | #golang generator | 4 | 
| [chmike/varint](https://github.com/chmike/varint) | variable length integer encoding using prefix code | 2 | 
| [rkoesters/xdg](https://github.com/rkoesters/xdg) | FreeDesktop.org (xdg) Specs implemented in Go | 28 | 
| [go-xkg/xkg](https://github.com/go-xkg/xkg) | User level X Keyboard Grabber | 53 | 
| [ulikunitz/xz](https://github.com/ulikunitz/xz) | Pure golang package for reading and writing xz-compressed files | 371 | 


# Natural Language Processing

| Repository | Description | Stars |
|:---|:---|:---:|


# Language Detection

| Repository | Description | Stars |
|:---|:---|:---:|


# Morphological Analyzers

| Repository | Description | Stars |
|:---|:---|:---:|


# Slugifiers

| Repository | Description | Stars |
|:---|:---|:---:|


# Tokenizers

| Repository | Description | Stars |
|:---|:---|:---:|


# Translation

| Repository | Description | Stars |
|:---|:---|:---:|


# Transliteration

| Repository | Description | Stars |
|:---|:---|:---:|


# Networking

| Repository | Description | Stars |
|:---|:---|:---:|
| [mdlayher/arp](https://github.com/mdlayher/arp) | Package arp implements the ARP protocol, as described in RFC 826. MIT Licensed. | 281 | 
| [StabbyCutyou/buffstreams](https://github.com/StabbyCutyou/buffstreams) | A library to simplify writing applications using TCP sockets to stream protobuff messages | 247 | 
| [zubairhamed/canopus](https://github.com/zubairhamed/canopus) | CoAP Client/Server implementing RFC 7252 for the Go Language | 148 | 
| [yl2chen/cidranger](https://github.com/yl2chen/cidranger) | Fast IP to CIDR lookup in Golang | 682 | 
| [mdlayher/dhcp6](https://github.com/mdlayher/dhcp6) | Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. MIT Licensed. | 73 | 
| [miekg/dns](https://github.com/miekg/dns) | DNS library in Go | 6174 | 
| [mosajjal/dnsmonster](https://github.com/mosajjal/dnsmonster) | Passive DNS Capture/Monitoring Framework | 169 | 
| [DarthPestilane/easytcp](https://github.com/DarthPestilane/easytcp) | :sparkles: :rocket: EasyTCP is a light-weight TCP framework written in Go (Golang), built with message router. EasyTCP helps you build a TCP server easily fast and less painful. | 290 | 
| [songgao/ether](https://github.com/songgao/ether) | A Go package for sending and receiving ethernet frames. Currently supporting Linux, Freebsd, and OS X. | 76 | 
| [mdlayher/ethernet](https://github.com/mdlayher/ethernet) | Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. MIT Licensed. | 240 | 
| [valyala/fasthttp](https://github.com/valyala/fasthttp) | Fast HTTP package for Go. Tuned for high performance. Zero memory allocations in hot paths. Up to 10x faster than net/http | 17398 | 
| [fortio/fortio](https://github.com/fortio/fortio) | Fortio load testing library, command line tool, advanced echo server and web UI in go (golang). Allows to specify a set query-per-second load and record latency histograms and other useful stats. | 2340 | 
| [jlaffaye/ftp](https://github.com/jlaffaye/ftp) | FTP client package for Go | 911 | 
| [fclairamb/ftpserverlib](https://github.com/fclairamb/ftpserverlib) | golang ftp server library | 290 | 
| [xtaci/gaio](https://github.com/xtaci/gaio) | High performance async-io(proactor) networking for Golang。golangのための高性能非同期io(proactor)ネットワーキング | 416 | 
| [Allenxuxu/gev](https://github.com/Allenxuxu/gev) | 🚀Gev is a lightweight, fast non-blocking TCP network library / websocket server based on Reactor mode. Support custom protocols to quickly and easily build high-performance servers.  | 1368 | 
| [jimlambrt/gldap](https://github.com/jimlambrt/gldap) | Build LDAP services w/ Go | 60 | 
| [DrmagicE/gmqtt](https://github.com/DrmagicE/gmqtt) | Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.x and V5 in golang | 597 | 
| [panjf2000/gnet](https://github.com/panjf2000/gnet) | 🚀 gnet is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go./ gnet 是一个高性能、轻量级、非阻塞的事件驱动 Go 网络框架。 | 6209 | 
| [google/gnxi](https://github.com/google/gnxi) | gNXI Tools - gRPC Network Management/Operations Interface Tools | 202 | 
| [hashicorp/go-getter](https://github.com/hashicorp/go-getter) | Package for downloading things from a string URL using a variety of protocols. | 1312 | 
| [joeig/go-powerdns](https://github.com/joeig/go-powerdns) | Go PowerDNS 4.x API Client | 57 | 
| [ccding/go-stun](https://github.com/ccding/go-stun) | A go implementation of the STUN client (RFC 3489 and RFC 5389) | 483 | 
| [osrg/gobgp](https://github.com/osrg/gobgp) | BGP implemented in the Go Programming Language | 2720 | 
| [averageflow/gohooks](https://github.com/averageflow/gohooks) | GoHooks make it easy to send and consume secured web-hooks from a Go application | 16 | 
| [sunwxg/golibwireshark](https://github.com/sunwxg/golibwireshark) |  | 24 | 
| [google/gopacket](https://github.com/google/gopacket) | Provides packet processing capabilities for Go | 4718 | 
| [akrennmair/gopcap](https://github.com/akrennmair/gopcap) | A simple wrapper around libpcap for the Go programming language | 438 | 
| [sunwxg/goshark](https://github.com/sunwxg/goshark) |  | 14 | 
| [gosnmp/gosnmp](https://github.com/gosnmp/gosnmp) | An SNMP library written in Go | 821 | 
| [gansidui/gotcp](https://github.com/gansidui/gotcp) | A Go package for quickly building tcp servers | 489 | 
| [cavaliergopher/grab](https://github.com/cavaliergopher/grab) | A download manager package for Go | 1009 | 
| [koofr/graval](https://github.com/koofr/graval) | An experimental go FTP server framework | 27 | 
| [qustavo/httplab](https://github.com/qustavo/httplab) | The interactive web server | 3786 | 
| [wzshiming/httpproxy](https://github.com/wzshiming/httpproxy) | HTTP proxy handler and dialer | 11 | 
| [c-robinson/iplib](https://github.com/c-robinson/iplib) | A library  for working with IP addresses and networks in Go | 74 | 
| [udhos/jazigo](https://github.com/udhos/jazigo) | Jazigo is a tool written in Go for retrieving configuration for multiple devices, similar to rancid, fetchconfig, oxidized, Sweet. | 174 | 
| [xtaci/kcp-go](https://github.com/xtaci/kcp-go) |  A Crypto-Secure, Production-Grade Reliable-UDP Library for golang with FEC  | 3259 | 
| [xtaci/kcptun](https://github.com/xtaci/kcptun) | A Stable & Secure Tunnel based on KCP with N:M multiplexing and FEC. Available for ARM, MIPS, 386 and AMD64。KCPプロトコルに基づく安全なトンネル。KCP 프로토콜을 기반으로 하는 보안 터널입니다。 | 12795 | 
| [fanux/lhttp](https://github.com/fanux/lhttp) | go websocket, a better way to buid your IM server | 664 | 
| [ian-kent/linkio](https://github.com/ian-kent/linkio) | Simulate network link speed | 50 | 
| [kirillDanshin/llb](https://github.com/kirillDanshin/llb) |  | 12 | 
| [hashicorp/mdns](https://github.com/hashicorp/mdns) | Simple mDNS client/server library in Golang | 852 | 
| [mqttPaho - The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.](https://eclipse.org/paho/clients/golang/?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [lesismal/nbio](https://github.com/lesismal/nbio) | Pure Go 1000k+ connections solution, support tls/http1.x/websocket and basically compatible with net/http, with high-performance and low memory cost, non-blocking, event-driven, easy-to-use. | 609 | 
| [cloudwego/netpoll](https://github.com/cloudwego/netpoll) | A high-performance non-blocking I/O networking framework, which focused on RPC scenarios, developed by ByteDance. | 2473 | 
| [intel-go/nff-go](https://github.com/intel-go/nff-go) | NFF-Go -Network Function Framework for GO (former YANFF) | 1174 | 
| [aerogo/packet](https://github.com/aerogo/packet) | :package: Send network packets over a TCP or UDP connection. | 66 | 
| [yahoo/panoptes-stream](https://github.com/yahoo/panoptes-stream) | A cloud native distributed streaming network telemetry. | 32 | 
| [schollz/peerdiscovery](https://github.com/schollz/peerdiscovery) | Pure-Go library for cross-platform local peer discovery using UDP multicast :woman: :repeat: :woman: | 537 | 
| [aybabtme/portproxy](https://github.com/aybabtme/portproxy) | TCP proxy, highjacks HTTP to allow CORS | 49 | 
| [polera/publicip](https://github.com/polera/publicip) | Go pkg for returning your public facing IP address. | 25 | 
| [lucas-clemente/quic-go](https://github.com/lucas-clemente/quic-go) | A QUIC implementation in pure go | 6425 | 
| [mdlayher/raw](https://github.com/mdlayher/raw) | Package raw enables reading and writing data at the device driver level for a network interface.  MIT Licensed. | 422 | 
| [pkg/sftp](https://github.com/pkg/sftp) | SFTP support for the go.crypto/ssh package | 1119 | 
| [gliderlabs/ssh](https://github.com/gliderlabs/ssh) | Easy SSH servers in Golang | 2409 | 
| [eduardonunesp/sslb](https://github.com/eduardonunesp/sslb) | Golang Super Simple Load Balance | 136 | 
| [gortc/stun](https://github.com/gortc/stun) | Fast RFC 5389 STUN implementation in go | 484 | 
| [firstrow/tcp_server](https://github.com/firstrow/tcp_server) | golang tcp server | 409 | 
| [two/tspool](https://github.com/two/tspool) | tcp server pool | 12 | 
| [anacrolix/utp](https://github.com/anacrolix/utp) | Use anacrolix/go-libutp instead | 162 | 
| [yahoo/vssh](https://github.com/yahoo/vssh) | Go Library to Execute Commands Over SSH at Scale | 788 | 
| [songgao/water](https://github.com/songgao/water) | A simple TUN/TAP library written in native Go. | 1388 | 
| [pion/webrtc](https://github.com/pion/webrtc) | Pure Go implementation of the WebRTC API | 8938 | 
| [masterzen/winrm](https://github.com/masterzen/winrm) | Command-line tool and library for Windows remote command execution in Go | 358 | 
| [xfxdev/xtcp](https://github.com/xfxdev/xtcp) | A TCP Server Framework with graceful shutdown, custom protocol. | 131 | 


# HTTP Clients

| Repository | Description | Stars |
|:---|:---|:---:|
| [h2non/gentleman](https://github.com/h2non/gentleman) | Plugin-driven, extensible HTTP client toolkit for Go | 938 | 
| [hashicorp/go-cleanhttp](https://github.com/hashicorp/go-cleanhttp) |  | 240 | 
| [bozd4g/go-http-client](https://github.com/bozd4g/go-http-client) | An enhanced http client for Golang | 35 | 
| [wenerme/go-req](https://github.com/wenerme/go-req) | Declarative golang HTTP client | 13 | 
| [hashicorp/go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) | Retryable HTTP client in Go | 1133 | 
| [levigross/grequests](https://github.com/levigross/grequests) | A Go "clone" of the great and famous Requests library | 1869 | 
| [gojek/heimdall](https://github.com/gojek/heimdall) | An enhanced HTTP client for Go | 2221 | 
| [valord577/httpc](https://github.com/valord577/httpc) | A customizable and simple HTTP client library. Only depend on the stdlib HTTP client. | 4 | 
| [ybbus/httpretry](https://github.com/ybbus/httpretry) | Enriches the standard go http client with retry functionality. | 18 | 
| [sethgrid/pester](https://github.com/sethgrid/pester) | Go (golang) http calls with retries and backoff  | 582 | 
| [imroc/req](https://github.com/imroc/req) | Simple Go HTTP client with Black Magic | 2199 | 
| [monaco-io/request](https://github.com/monaco-io/request) | go request, go http client | 190 | 
| [carlmjohnson/requests](https://github.com/carlmjohnson/requests) | HTTP requests for Gophers | 297 | 
| [go-resty/resty](https://github.com/go-resty/resty) | Simple HTTP and REST client library for Go | 5812 | 
| [ddo/rq](https://github.com/ddo/rq) | A nicer interface for golang stdlib HTTP client | 40 | 
| [dghubble/sling](https://github.com/dghubble/sling) | A Go HTTP client library for creating and sending API requests | 1423 | 


# OpenGL

| Repository | Description | Stars |
|:---|:---|:---:|
| [go-gl/gl](https://github.com/go-gl/gl) | Go bindings for OpenGL (generated via glow) | 886 | 
| [go-gl/glfw](https://github.com/go-gl/glfw) | Go bindings for GLFW 3 | 1228 | 
| [technohippy/go-glmatrix](https://github.com/technohippy/go-glmatrix) | go-glmatrix is a golang version of glMatrix, which is "designed to perform vector and matrix operations stupidly fast". | 3 | 
| [goxjs/gl](https://github.com/goxjs/gl) | Go cross-platform OpenGL bindings. | 155 | 
| [goxjs/glfw](https://github.com/goxjs/glfw) | Go cross-platform glfw library for creating an OpenGL context and receiving events. | 74 | 
| [go-gl/mathgl](https://github.com/go-gl/mathgl) | A pure Go 3D math library. | 424 | 


# ORM

| Repository | Description | Stars |
|:---|:---|:---:|
| [Yiling-J/cacheme-go](https://github.com/Yiling-J/cacheme-go) | 🚀 Schema based, typed Redis caching/memoize framework for Go | 19 | 
| [ent/ent](https://github.com/ent/ent) | An entity framework for Go | 10242 | 
| [jschoedt/go-firestorm](https://github.com/jschoedt/go-firestorm) | Simple Go ORM for Google/Firebase Cloud Firestore | 27 | 
| [go-pg/pg](https://github.com/go-pg/pg) | Golang ORM with focus on PostgreSQL features and performance | 5008 | 
| [jirfag/go-queryset](https://github.com/jirfag/go-queryset) | 100% type-safe ORM for Go (Golang) with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support. GORM under the hood. | 656 | 
| [rushteam/gosql](https://github.com/rushteam/gosql) | golang orm and sql builder | 158 | 
| [huandu/go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) | A flexible and powerful SQL string builder library plus a zero-config ORM. | 743 | 
| [gosuri/go-store](https://github.com/gosuri/go-store) | A simple and fast Redis backed key-value store library for Go | 107 | 
| [go-gorm/gorm](https://github.com/go-gorm/gorm) | The fantastic ORM library for Golang, aims to be developer friendly | 27318 | 
| [xxjwxc/gormt](https://github.com/xxjwxc/gormt) | database to golang struct | 1770 | 
| [go-gorp/gorp](https://github.com/go-gorp/gorp) | Go Relational Persistence - an ORM-ish library for Go | 3556 | 
| [Fs02/grimoire](https://github.com/Fs02/grimoire) | Database access layer for golang | 156 | 
| [abrahambotros/lore](https://github.com/abrahambotros/lore) | Light Object-Relational Environment (LORE) provides a simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go | 10 | 
| [marlow/marlow](https://github.com/marlow/marlow) | persistence layer code generation for golang | 11 | 
| [gobuffalo/pop](https://github.com/gobuffalo/pop) | A Tasty Treat For All Your Database Needs | 1184 | 
| [prisma/prisma-client-go](https://github.com/prisma/prisma-client-go) | Prisma Client Go is an auto-generated and fully type-safe database client | 1172 | 
| [go-reform/reform](https://github.com/go-reform/reform) | A better ORM for Go, based on non-empty interfaces and code generation. | 1232 | 
| [go-rel/rel](https://github.com/go-rel/rel) | :gem: Modern ORM for Golang - Testable, Extendable and Crafted Into a Clean and Elegant API | 497 | 
| [volatiletech/sqlboiler](https://github.com/volatiletech/sqlboiler) | Generate a Go ORM tailored to your database schema. | 4734 | 
| [upper/db](https://github.com/upper/db) | Data access layer for PostgreSQL, CockroachDB, MySQL, SQLite and MongoDB with ORM-like features. | 2918 | 
| [XORM - Simple and powerful ORM for Go. (Support: MySQL, MyMysql, PostgreSQL, Tidb, SQLite3, MsSql and Oracle).](https://gitea.com/xorm/xorm?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [albrow/zoom](https://github.com/albrow/zoom) | A blazing-fast datastore and querying engine for Go built on Redis. | 286 | 


# Package Management

| Repository | Description | Stars |
|:---|:---|:---:|
| [go modules - Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 


# Performance

| Repository | Description | Stars |
|:---|:---|:---:|


# Query Language

| Repository | Description | Stars |
|:---|:---|:---:|


# Resource Embedding

| Repository | Description | Stars |
|:---|:---|:---:|


# Science and Data Analysis

| Repository | Description | Stars |
|:---|:---|:---:|
| [ndabAP/assocentity](https://github.com/ndabAP/assocentity) | Package assocentity returns the average distance from words to a given entity | 8 | 
| [seanhagen/bradleyterry](https://github.com/seanhagen/bradleyterry) | Package to do Bradley-Terry Model pairwise compairsons | 5 | 
| [nikolaydubina/calendarheatmap](https://github.com/nikolaydubina/calendarheatmap) | 📅 Calendar heatmap inspired by GitHub contribution activity  | 342 | 
| [vdobler/chart](https://github.com/vdobler/chart) | Provide basic charts in go | 715 | 
| [rocketlaunchr/dataframe-go](https://github.com/rocketlaunchr/dataframe-go) | DataFrames for Go: For statistics, machine-learning, and data manipulation/exploration | 768 | 
| [db47h/decimal](https://github.com/db47h/decimal) | An arbitrary-precision decimal floating-point arithmetic package for Go | 24 | 
| [soniah/evaler](https://github.com/soniah/evaler) | Implements a simple floating point arithmetic expression evaluator in Go (golang). | 47 | 
| [VividCortex/ewma](https://github.com/VividCortex/ewma) | Exponentially Weighted Moving Average algorithms for Go. | 366 | 
| [skelterjohn/geom](https://github.com/skelterjohn/geom) | 2d geometry for golang | 50 | 
| [mjibson/go-dsp](https://github.com/mjibson/go-dsp) | Digital Signal Processing for Go | 770 | 
| [milosgajdos/go-estimate](https://github.com/milosgajdos/go-estimate) | State estimation and filtering algorithms in Go | 93 | 
| [ThePaw/go-gt](https://github.com/ThePaw/go-gt) | Automatically exported from code.google.com/p/go-gt | 6 | 
| [soypat/godesim](https://github.com/soypat/godesim) | ODE system solver made simple. For IVPs (initial value problems). | 19 | 
| [kzahedi/goent](https://github.com/kzahedi/goent) | GO Implementation of Entropy Measures | 26 | 
| [VividCortex/gohistogram](https://github.com/VividCortex/gohistogram) | Streaming approximate histograms in Go | 161 | 
| [gonum/gonum](https://github.com/gonum/gonum) | Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more | 5624 | 
| [gonum/plot](https://github.com/gonum/plot) | A repository for plotting and visualizing data | 2134 | 
| [gyuho/goraph](https://github.com/gyuho/goraph) | Package goraph implements graph data structure and algorithms. | 672 | 
| [cpmech/gosl](https://github.com/cpmech/gosl) | Linear algebra, eigenvalues, FFT, Bessel, elliptic, orthogonal polys, geometry, NURBS, numerical quadrature, 3D transfinite interpolation, random numbers, Mersenne twister, probability distributions, optimisation, differential equations. | 1647 | 
| [OGFris/GoStats](https://github.com/OGFris/GoStats) | GoStats is a go library for math statistics mostly used in ML domains, it covers most of the statistical measures functions. | 20 | 
| [yourbasic/graph](https://github.com/yourbasic/graph) | Graph algorithms and data structures | 535 | 
| [nikolaydubina/jsonl-graph](https://github.com/nikolaydubina/jsonl-graph) | 🏝 JSONL Graph Tools | 58 | 
| [ChristopherRabotin/ode](https://github.com/ChristopherRabotin/ode) | An ordinary differential equation solving library in golang. | 17 | 
| [paulmach/orb](https://github.com/paulmach/orb) | Types and utilities for working with 2d geometry in Golang | 497 | 
| [alixaxel/pagerank](https://github.com/alixaxel/pagerank) | Weighted PageRank implementation in Go | 74 | 
| [sgreben/piecewiselinear](https://github.com/sgreben/piecewiselinear) | tiny linear interpolation library for go (factored out from https://github.com/sgreben/yeetgif) | 22 | 
| [claygod/PiHex](https://github.com/claygod/PiHex) | PiHex Library, written in Go, generates a hexadecimal number sequence in the number Pi in the range from 0 to 10,000,000. | 15 | 
| [khezen/rootfinding](https://github.com/khezen/rootfinding) | root-finding library | 7 | 
| [james-bowman/sparse](https://github.com/james-bowman/sparse) | Sparse matrix formats for linear algebra supporting scientific and machine learning applications | 130 | 
| [montanaflynn/stats](https://github.com/montanaflynn/stats) | A well tested and comprehensive Golang statistics library package with no dependencies. | 2318 | 
| [nytlabs/streamtools](https://github.com/nytlabs/streamtools) | tools for working with streams of data | 1317 | 
| [DavidBelicza/TextRank](https://github.com/DavidBelicza/TextRank) | :wink: :cyclone: :strawberry: TextRank implementation in Golang with extendable features (summarization, phrase extraction) and multithreading (goroutine). | 151 | 
| [tchayen/triangolatte](https://github.com/tchayen/triangolatte) | 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. | 26 | 


# Security

| Repository | Description | Stars |
|:---|:---|:---:|
| [hlandau/acmetool](https://github.com/hlandau/acmetool) | :lock: acmetool, an automatic certificate acquisition tool for ACME (Let's Encrypt) | 1897 | 
| [cossacklabs/acra](https://github.com/cossacklabs/acra) | Database security suite. Database proxy with field-level encryption, search through encrypted data, SQL injections prevention, intrusion detection, honeypots. Supports client-side and proxy-side ("transparent") encryption. SQL, NoSQL. | 958 | 
| [FiloSottile/age](https://github.com/FiloSottile/age) | A simple, modern and secure encryption tool (and Go library) with small explicit keys, no config options, and UNIX-style composability. | 10181 | 
| [andskur/argon2-hashing](https://github.com/andskur/argon2-hashing) | A light package for generating and comparing password hashing with argon2 in Go | 16 | 
| [raja/argon2pw](https://github.com/raja/argon2pw) | Argon2 password hashing package for go with constant time hash comparison | 89 | 
| [autocert - Auto provision Let’s Encrypt certificates and start a TLS server.](https://godoc.org/golang.org/x/crypto/acme/autocert?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [jaredfolkins/badactor](https://github.com/jaredfolkins/badactor) | BadActor.org An in-memory application driven jailer written in Go | 305 | 
| [Ullaakut/cameradar](https://github.com/Ullaakut/cameradar) | Cameradar hacks its way into RTSP videosurveillance cameras | 2867 | 
| [mvmaasakkers/certificates](https://github.com/mvmaasakkers/certificates) | An opinionated helper for generating tls certificates | 23 | 
| [caddyserver/certmagic](https://github.com/caddyserver/certmagic) | Automatic HTTPS for any Go program: fully-managed TLS certificate issuance and renewal | 3973 | 
| [golang-module/dongle](https://github.com/golang-module/dongle) | A simple, semantic and developer-friendly golang package for encoding&decoding and encryption&decryption | 76 | 
| [prashantgupta24/firewalld-rest](https://github.com/prashantgupta24/firewalld-rest) | A rest application to update firewalld rules on a linux server  | 316 | 
| [m1/go-generate-password](https://github.com/m1/go-generate-password) | Password generator written in Go | 31 | 
| [tg123/go-htpasswd](https://github.com/tg123/go-htpasswd) | Apache htpasswd Parser for Go. | 23 | 
| [wagslane/go-password-validator](https://github.com/wagslane/go-password-validator) | Validate the Strength of a Password in Go | 326 | 
| [hillu/go-yara](https://github.com/hillu/go-yara) | Go bindings for YARA | 247 | 
| [dwin/goArgonPass](https://github.com/dwin/goArgonPass) | goArgonPass is a Argon2 Password utility package for Go using the crypto library package Argon2 designed to be compatible with Passlib for Python and Argon2 PHP. Argon2 was the winner of the most recent Password Hashing Competition. This is designed for use anywhere password hashing and verification might be needed and is intended to replace implementations using bcrypt or Scrypt. | 14 | 
| [dwin/goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) | A probably paranoid Golang utility library for securely hashing and encrypting passwords based on the Dropbox method. This implementation uses Blake2b, Scrypt and XSalsa20-Poly1305 (via NaCl SecretBox) to create secure password hashes that are also encrypted using a master passphrase. | 47 | 
| [Interpol - Rule-based data generator for fuzzing and penetration testing.](https://bitbucket.org/vahidi/interpol?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [go-acme/lego](https://github.com/go-acme/lego) | Let's Encrypt/ACME client and library written in Go | 5171 | 
| [awnumar/memguard](https://github.com/awnumar/memguard) | Secure software enclave for storage of sensitive information in memory. | 2142 | 
| [kevinburke/nacl](https://github.com/kevinburke/nacl) | Pure Go implementation of the NaCL set of API's | 521 | 
| [pjebs/optimus-go](https://github.com/pjebs/optimus-go) | ID hashing and Obfuscation using Knuth's Algorithm | 305 | 
| [hlandau/passlib](https://github.com/hlandau/passlib) | :key: Idiotproof golang password validation library inspired by Python's passlib | 263 | 
| [rsjethani/secret](https://github.com/rsjethani/secret) | Prevent your secrets from leaking into logs, std* etc. | 7 | 
| [unrolled/secure](https://github.com/unrolled/secure) | HTTP middleware for Go that facilitates some quick security wins. | 1897 | 
| [xaionaro-go/secureio](https://github.com/xaionaro-go/secureio) | An easy-to-use XChaCha20-encryption wrapper for io.ReadWriteCloser (even lossy UDP) using ECDH key exchange algorithm, ED25519 signatures and Blake3+Poly1305 checksums/message-authentication for Go (golang). Also a multiplexer. | 24 | 
| [elithrar/simple-scrypt](https://github.com/elithrar/simple-scrypt) | A convenience library for generating, comparing and inspecting password hashes using the scrypt KDF in Go 🔑 | 176 | 
| [ssh-vault/ssh-vault](https://github.com/ssh-vault/ssh-vault) | 🌰  encrypt/decrypt using ssh keys | 342 | 
| [adrianosela/sslmgr](https://github.com/adrianosela/sslmgr) | A layer of abstraction the around acme/autocert certificate manager (Golang) | 13 | 
| [cossacklabs/themis](https://github.com/cossacklabs/themis) | Easy to use cryptographic framework for data protection: secure messaging with forward secrecy and secure data storage. Has unified APIs across 14 platforms. | 1484 | 


# Serialization

| Repository | Description | Stars |
|:---|:---|:---:|
| [Logicalis/asn1](https://github.com/Logicalis/asn1) | Asn.1 BER and DER encoding library for golang. | 48 | 
| [glycerine/bambam](https://github.com/glycerine/bambam) | auto-generate capnproto schema from your golang source files. Depends on go-capnproto-1.0 at https://github.com/glycerine/go-capnproto | 64 | 
| [csweichel/bel](https://github.com/csweichel/bel) | Generate TypeScript interfaces from Go structs/interfaces - useful for JSON RPC | 20 | 
| [ghostiam/binstruct](https://github.com/ghostiam/binstruct) | Golang binary decoder for mapping data into the structure | 44 | 
| [fxamacker/cbor](https://github.com/fxamacker/cbor) | CBOR codec (RFC 8949) with CBOR tags, Go struct tags (toarray, keyasint, omitempty), float64/32/16, big.Int, and fuzz tested billions of execs.  | 386 | 
| [pascaldekloe/colfer](https://github.com/pascaldekloe/colfer) | binary serialization format | 653 | 
| [jszwec/csvutil](https://github.com/jszwec/csvutil) | csvutil provides fast and idiomatic mapping between CSV and Go (golang) values. | 684 | 
| [epiclabs-io/elastic](https://github.com/epiclabs-io/elastic) | Converts go types no matter what | 16 | 
| [huydang284/fixedwidth](https://github.com/huydang284/fixedwidth) | A Go package for encode/decode fixed-width data | 6 | 
| [o1egl/fwencoder](https://github.com/o1egl/fwencoder) | Fixed width file parser (encoder/decoder) in GO (golang) | 20 | 
| [glycerine/go-capnproto](https://github.com/glycerine/go-capnproto) | Cap'n Proto library and parser for go. This is go-capnproto-1.0, and does not have rpc. See https://github.com/zombiezen/go-capnproto2 for 2.0 which has rpc and capabilities. | 280 | 
| [ugorji/go](https://github.com/ugorji/go) | idiomatic codec and rpc lib for msgpack, cbor, json, etc. msgpack.org[Go] | 1643 | 
| [sbourlon/go-lctree](https://github.com/sbourlon/go-lctree) | go-lctree provides a CLI and Go primitives to serialize and deserialize LeetCode binary trees (e.g. "[5,4,7,3,null,2,null,-1,null,9]"). | 3 | 
| [gogo/protobuf](https://github.com/gogo/protobuf) | [Looking for new ownership] Protocol Buffers for Go with Gadgets | 5155 | 
| [golang/protobuf](https://github.com/golang/protobuf) | Go support for Google's protocol buffers | 8316 | 
| [json-iterator/go](https://github.com/json-iterator/go) | A high-performance 100% compatible drop-in replacement of "encoding/json" | 10617 | 
| [mitchellh/mapstructure](https://github.com/mitchellh/mapstructure) | Go library for decoding generic map values into native Go structures and vice versa. | 5508 | 
| [yvasiyarov/php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) | PHP session encoder/decoder written in Go | 154 | 
| [vimeda/pletter](https://github.com/vimeda/pletter) | A standard way to wrap a proto message | 17 | 
| [danhper/structomap](https://github.com/danhper/structomap) | Easily and dynamically generate maps from Go static structures | 131 | 
| [recolude/unitpacking](https://github.com/recolude/unitpacking) | A library for storing unit vectors in a representation that lends itself to saving space on disk. | 4 | 


# Server Applications

| Repository | Description | Stars |
|:---|:---|:---:|


# Stream Processing

| Repository | Description | Stars |
|:---|:---|:---:|
| [reugn/go-streams](https://github.com/reugn/go-streams) | A lightweight stream processing library for Go | 878 | 
| [whitaker-io/machine](https://github.com/whitaker-io/machine) | Machine is a workflow/pipeline library for processing data | 109 | 
| [youthlin/stream](https://github.com/youthlin/stream) | Go Stream, like Java 8 Stream. | 53 | 


# Template Engines

| Repository | Description | Stars |
|:---|:---|:---:|
| [benbjohnson/ego](https://github.com/benbjohnson/ego) | An ERB-style templating language for Go. | 515 | 
| [dannyvankooten/extemplate](https://github.com/dannyvankooten/extemplate) | Wrapper package for Go's template/html to allow for easy file-based template inheritance. | 43 | 
| [valyala/fasttemplate](https://github.com/valyala/fasttemplate) | Simple and fast template engine for Go | 576 | 
| [m1/gospin](https://github.com/m1/gospin) | Article spinning and spintax/spinning syntax engine written in Go, useful for A/B, testing pieces of text/articles and creating more natural conversations | 35 | 
| [goradd/got](https://github.com/goradd/got) | GoT is a template engine that turns templates into Go code to compile into your app. | 2 | 
| [foolin/goview](https://github.com/foolin/goview) | Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. | 272 | 
| [CloudyKit/jet](https://github.com/CloudyKit/jet) | Jet  template engine | 903 | 
| [osteele/liquid](https://github.com/osteele/liquid) | A Liquid template engine in Go | 148 | 
| [johnfercher/maroto](https://github.com/johnfercher/maroto) | A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. | 578 | 
| [flosch/pongo2](https://github.com/flosch/pongo2) | Django-syntax like template-engine for Go | 2190 | 
| [valyala/quicktemplate](https://github.com/valyala/quicktemplate) | Fast, powerful, yet easy to use template engine for Go. Optimized for speed, zero memory allocations in hot paths. Up to 20x faster than html/template | 2366 | 
| [aymerick/raymond](https://github.com/aymerick/raymond) | Handlebars for golang | 466 | 
| [sipin/gorazor](https://github.com/sipin/gorazor) | Razor view engine for go | 796 | 
| [robfig/soy](https://github.com/robfig/soy) | Go implementation for Soy templates (Google Closure templates) | 160 | 
| [Masterminds/sprig](https://github.com/Masterminds/sprig) | Useful template functions for Go templates. | 2868 | 
| [lucasepe/tbd](https://github.com/lucasepe/tbd) | "to be defined" - a really simple way to create text templates with placeholders | 16 | 


# Testing

| Repository | Description | Stars |
|:---|:---|:---:|
| [Testing Frameworks apitest - Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams. assert - Basic Assertion Library used along side native go testing, with building blocks for custom assertions. badio - Extensions to Go’s testing/iotest package. baloo - Expressive and versatile end-to-end HTTP API testing made easy. biff - Bifurcation testing framework, BDD compatible. charlatan - Tool to generate fake interface implementations for tests. commander - Tool for testing cli applications on windows, linux and osx. covergates - Self-hosted code coverage report review and management service. cupaloy - Simple snapshot testing addon for your test framework. dbcleaner - Clean database for testing purpose, inspired by database_cleaner in Ruby. dsunit - Datastore testing for SQL, NoSQL, structured files. embedded-postgres - Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test. endly - Declarative end to end functional testing. fixenv - Fixture manage engine, inspired by pytest fixtures. flute - HTTP client testing framework. frisby - REST API testing framework. gherkingen - BDD boilerplate generator and framework. ginkgo - BDD Testing Framework for Go. gnomock - integration testing with real dependencies (database, cache, even Kubernetes or AWS) running in Docker, without mocks. go-carpet - Tool for viewing test coverage in terminal. go-cmp - Package for comparing Go values in tests. go-hit - Hit is an http integration test framework written in golang. go-mutesting - Mutation testing for Go source code. go-testdeep - Extremely flexible golang deep comparison, extends the go testing package. go-testpredicate - Test predicate style assertions library with extensive diagnostics output. go-vcr - Record and replay your HTTP interactions for fast, deterministic and accurate tests. goblin - Mocha like testing framework fo Go. goc - Goc is a comprehensive coverage testing system for The Go Programming Language. gocheck - More advanced testing framework alternative to gotest. GoConvey - BDD-style framework with web UI and live reload. gocrest - Composable hamcrest-like matchers for Go assertions. godog - Cucumber or Behat like BDD framework for Go. gofight - API Handler Testing for Golang Router framework. gogiven - YATSPEC-like BDD testing framework for Go. gomatch - library created for testing JSON against patterns. gomega - Rspec like matcher/assertion library. GoSpec - BDD-style testing framework for the Go programming language. gospecify - This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec. gosuite - Brings lightweight test suites with setup/teardown facilities to testing by leveraging Go1.7’s Subtests. gotest.tools - A collection of packages to augment the go testing package and support common patterns. Hamcrest - fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results. httpexpect - Concise, declarative, and easy to use end-to-end HTTP and REST API testing. is - Professional lightweight testing mini-framework for Go. jsonassert - Package for verifying that your JSON payloads are serialized correctly. omg.testingtools - The simple library for change a values of private fields for testing. restit - Go micro framework to help writing RESTful API integration test. schema - Quick and easy expression matching for JSON schemas used in requests and responses. stop-and-go - Testing helper for concurrency. testcase - Idiomatic testing framework for Behavior Driven Development. testfixtures - A helper for Rails’ like test fixtures to test database applications. Testify - Sacred extension to the standard go testing package. testmd - Convert markdown snippets into testable go code. testsql - Generate test data from SQL files before testing and clear it after finished. testza - Full-featured test framework with nice colorized output. trial - Quick and easy extendable assertions without introducing much boilerplate. Tt - Simple and colorful test tools. wstest - Websocket client for unit-testing a websocket http.Handler.](https://apitest.dev?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [apitest - Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.](https://apitest.dev?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [go-playground/assert](https://github.com/go-playground/assert) | :exclamation:Basic Assertion Library used along side native go testing, with building blocks for custom assertions | 38 | 
| [cavaliergopher/badio](https://github.com/cavaliergopher/badio) | Extensions to Go's testing/iotest package | 9 | 
| [h2non/baloo](https://github.com/h2non/baloo) | Expressive end-to-end HTTP API testing made easy in Go | 722 | 
| [fulldump/biff](https://github.com/fulldump/biff) | Bifurcation Framework for testing and use cases | 10 | 
| [percolate/charlatan](https://github.com/percolate/charlatan) | Go Interface Mocking Tool | 194 | 
| [commander-cli/commander](https://github.com/commander-cli/commander) | Test your command line interfaces on windows, linux and osx and nodes viá ssh and docker | 197 | 
| [covergates/covergates](https://github.com/covergates/covergates) | The portal gates to coverage reports | 47 | 
| [bradleyjkemp/cupaloy](https://github.com/bradleyjkemp/cupaloy) | Simple Go snapshot testing | 205 | 
| [khaiql/dbcleaner](https://github.com/khaiql/dbcleaner) | Clean database for testing, inspired by database_cleaner for Ruby | 136 | 
| [viant/dsunit](https://github.com/viant/dsunit) | Datastore Testibility | 39 | 
| [fergusstrange/embedded-postgres](https://github.com/fergusstrange/embedded-postgres) | Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test | 349 | 
| [viant/endly](https://github.com/viant/endly) | End to end functional test and automation framework | 204 | 
| [rekby/fixenv](https://github.com/rekby/fixenv) |  | 5 | 
| [suzuki-shunsuke/flute](https://github.com/suzuki-shunsuke/flute) | Golang HTTP client testing framework | 16 | 
| [hofstadter-io/frisby](https://github.com/hofstadter-io/frisby) | API testing framework inspired by frisby-js | 271 | 
| [hedhyw/gherkingen](https://github.com/hedhyw/gherkingen) | Behaviour Driven Development tests generator for Golang | 39 | 
| [ginkgo - BDD Testing Framework for Go.](https://onsi.github.io/ginkgo/?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [orlangure/gnomock](https://github.com/orlangure/gnomock) | Test your code without writing mocks with ephemeral Docker containers 📦 Setup popular services with just a couple lines of code ⏱️ No bash, no yaml, only code 💻 | 691 | 
| [msoap/go-carpet](https://github.com/msoap/go-carpet) | go-carpet - show test coverage in terminal for Go source files | 220 | 
| [google/go-cmp](https://github.com/google/go-cmp) | Package for comparing Go values in tests | 2768 | 
| [Eun/go-hit](https://github.com/Eun/go-hit) | http integration test framework | 73 | 
| [zimmski/go-mutesting](https://github.com/zimmski/go-mutesting) | Mutation testing for Go source code | 516 | 
| [maxatome/go-testdeep](https://github.com/maxatome/go-testdeep) | Extremely flexible golang deep comparison, extends the go testing package, tests HTTP APIs and provides tests suite | 275 | 
| [maargenton/go-testpredicate](https://github.com/maargenton/go-testpredicate) | Unit-testing predicates for Go. | 4 | 
| [dnaeon/go-vcr](https://github.com/dnaeon/go-vcr) | Record and replay your HTTP interactions for fast, deterministic and accurate tests | 866 | 
| [franela/goblin](https://github.com/franela/goblin) | Minimal and Beautiful Go testing framework | 842 | 
| [qiniu/goc](https://github.com/qiniu/goc) | A Comprehensive Coverage Testing System for The Go Programming Language | 498 | 
| [gocheck - More advanced testing framework alternative to gotest.](https://labix.org/gocheck?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [smartystreets/goconvey](https://github.com/smartystreets/goconvey) | Go testing in the browser. Integrates with `go test`. Write behavioral tests in Go. | 7068 | 
| [corbym/gocrest](https://github.com/corbym/gocrest) | GoCrest - Hamcrest-like matchers for Go | 81 | 
| [cucumber/godog](https://github.com/cucumber/godog) | Cucumber for golang | 1612 | 
| [appleboy/gofight](https://github.com/appleboy/gofight) | Testing API Handler written in Golang. | 399 | 
| [corbym/gogiven](https://github.com/corbym/gogiven) | gogiven - BDD testing framework for go that generates readable output directly from source code | 11 | 
| [jfilipczyk/gomatch](https://github.com/jfilipczyk/gomatch) | Library created for testing JSON against patterns. | 41 | 
| [gomega - Rspec like matcher/assertion library.](https://onsi.github.io/gomega/?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [luontola/gospec](https://github.com/luontola/gospec) | Testing framework for Go. Allows writing self-documenting tests/specifications, and executes them concurrently and safely isolated. [UNMAINTAINED] | 113 | 
| [stesla/gospecify](https://github.com/stesla/gospecify) | A BDD library for Go | 52 | 
| [pavlo/gosuite](https://github.com/pavlo/gosuite) | Test suites support for standard Go1.7 "testing" by leveraging Subtests feature | 10 | 
| [gotestyourself/gotest.tools](https://github.com/gotestyourself/gotest.tools) | A collection of packages to augment the go testing package and support common patterns. | 291 | 
| [rdrdr/hamcrest](https://github.com/rdrdr/hamcrest) | Hamcrest matchers for the Go programming language | 27 | 
| [gavv/httpexpect](https://github.com/gavv/httpexpect) | End-to-end HTTP and REST API testing for Go. | 1872 | 
| [matryer/is](https://github.com/matryer/is) | Professional lightweight testing mini-framework for Go. | 1377 | 
| [kinbiko/jsonassert](https://github.com/kinbiko/jsonassert) | A Go test assertion library for verifying that two representations of JSON are semantically equal | 77 | 
| [dedalqq/omg.testingtools](https://github.com/dedalqq/omg.testingtools) | This tool can be useful for writing a tests. If you want change private field in struct from imported libraries than it can help you. | 0 | 
| [go-restit/restit](https://github.com/go-restit/restit) | A Go library help testing your RESTful API application | 55 | 
| [jgroeneveld/schema](https://github.com/jgroeneveld/schema) | Quick and easy expression matching for JSON schemas used in requests and responses | 16 | 
| [elgohr/stop-and-go](https://github.com/elgohr/stop-and-go) | Testing helper for concurrency | 5 | 
| [adamluzsi/testcase](https://github.com/adamluzsi/testcase) | testcase is an opinionated testing framework based on BDD principles. | 82 | 
| [go-testfixtures/testfixtures](https://github.com/go-testfixtures/testfixtures) | Ruby on Rails like test fixtures for Go. Write tests against a real database | 756 | 
| [stretchr/testify](https://github.com/stretchr/testify) | A toolkit with common assertions and mocks that plays nicely with the standard library | 15794 | 
| [testmd - Convert markdown snippets into testable go code.](https://godoc.org/github.com/tvastar/test/cmd/testmd?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [zhulongcheng/testsql](https://github.com/zhulongcheng/testsql) | Generate test data from SQL files before testing and clear it after finished. | 12 | 
| [MarvinJWendt/testza](https://github.com/MarvinJWendt/testza) | Full-featured test framework for Go! Assertions, fuzzing, input testing, output capturing, and much more! 🍕 | 325 | 
| [jgroeneveld/trial](https://github.com/jgroeneveld/trial) | A simple assertion library for go | 5 | 
| [vcaesar/tt](https://github.com/vcaesar/tt) | Simple and colorful test tools | 4 | 
| [posener/wstest](https://github.com/posener/wstest) | go websocket client for unit testing of a websocket handler | 87 | 
| [maxbrunsfeld/counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) | A tool for generating self-contained, type-safe test doubles in go | 638 | 
| [maxbrunsfeld/counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) | A tool for generating self-contained, type-safe test doubles in go | 638 | 
| [genmock - Go mocking system with code generator for building calls of the interface methods.](https://gitlab.com/so_literate/genmock?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [elgohr/go-localstack](https://github.com/elgohr/go-localstack) | Go Wrapper for using localstack | 41 | 
| [DATA-DOG/go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) | Sql mock driver for golang to test database interactions | 4251 | 
| [DATA-DOG/go-txdb](https://github.com/DATA-DOG/go-txdb) | Immutable transaction isolated sql driver for golang | 427 | 
| [h2non/gock](https://github.com/h2non/gock) | HTTP traffic mocking and testing made easy in Go ༼ʘ̚ل͜ʘ̚༽ | 1593 | 
| [golang/mock](https://github.com/golang/mock) | GoMock is a mocking framework for the Go programming language. | 6944 | 
| [seborama/govcr](https://github.com/seborama/govcr) | HTTP mock for Golang: record and replay HTTP/HTTPS interactions for offline testing | 101 | 
| [SpectoLabs/hoverfly](https://github.com/SpectoLabs/hoverfly) | Lightweight service virtualization/API simulation tool for developers and testers | 1857 | 
| [jarcoal/httpmock](https://github.com/jarcoal/httpmock) | HTTP mocking for Golang | 1328 | 
| [gojuno/minimock](https://github.com/gojuno/minimock) | Powerful mock generation tool for Go programming language | 450 | 
| [vektra/mockery](https://github.com/vektra/mockery) | A mock code autogenerator for Golang | 3375 | 
| [tv42/mockhttp](https://github.com/tv42/mockhttp) | Mock object for Go http.ResponseWriter | 21 | 
| [pasdam/mockit](https://github.com/pasdam/mockit) | Library that make mocking of Go functions/methods easy | 8 | 
| [cabify/timex](https://github.com/cabify/timex) | A test-friendly replacement for golang's time package | 60 | 
| [dvyukov/go-fuzz](https://github.com/dvyukov/go-fuzz) | Randomized testing for Go | 4356 | 
| [dvyukov/go-fuzz](https://github.com/dvyukov/go-fuzz) | Randomized testing for Go | 4356 | 
| [google/gofuzz](https://github.com/google/gofuzz) | Fuzz testing for go. | 1245 | 
| [zimmski/tavor](https://github.com/zimmski/tavor) | A generic fuzzing and delta-debugging framework | 232 | 
| [mafredri/cdp](https://github.com/mafredri/cdp) | Package cdp provides type-safe bindings for the Chrome DevTools Protocol (CDP), written in the Go programming language. | 599 | 
| [mafredri/cdp](https://github.com/mafredri/cdp) | Package cdp provides type-safe bindings for the Chrome DevTools Protocol (CDP), written in the Go programming language. | 599 | 
| [chromedp/chromedp](https://github.com/chromedp/chromedp) | A faster, simpler way to drive browsers supporting the Chrome DevTools Protocol. | 7387 | 
| [aerokube/ggr](https://github.com/aerokube/ggr) | A lightweight load balancer used to create big Selenium clusters | 283 | 
| [playwright-community/playwright-go](https://github.com/playwright-community/playwright-go) | Playwright for Go a browser automation library to control Chromium, Firefox and WebKit with a single API. | 716 | 
| [go-rod/rod](https://github.com/go-rod/rod) | A Devtools driver for web automation and scraping | 2256 | 
| [aerokube/selenoid](https://github.com/aerokube/selenoid) | Selenium Hub successor running browsers within containers. Scalable, immutable, self hosted Selenium-Grid on any platform with single binary. | 2121 | 
| [pingcap/failpoint](https://github.com/pingcap/failpoint) | An implementation of failpoints for Golang. | 658 | 
| [pingcap/failpoint](https://github.com/pingcap/failpoint) | An implementation of failpoints for Golang. | 658 | 


# Text Processing

| Repository | Description | Stars |
|:---|:---|:---:|


# Formatters

| Repository | Description | Stars |
|:---|:---|:---:|


# Markup Languages

| Repository | Description | Stars |
|:---|:---|:---:|


# Regular Expressions

| Repository | Description | Stars |
|:---|:---|:---:|


# Sanitation

| Repository | Description | Stars |
|:---|:---|:---:|


# Scrapers

| Repository | Description | Stars |
|:---|:---|:---:|


# RSS

| Repository | Description | Stars |
|:---|:---|:---:|


# Third-party APIs

| Repository | Description | Stars |
|:---|:---|:---:|
| [mehanizm/airtable](https://github.com/mehanizm/airtable) | Simple golang airtable API wrapper | 37 | 
| [ngs/go-amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) | Go Client Library for Amazon Product Advertising API | 51 | 
| [ChimeraCoder/anaconda](https://github.com/ChimeraCoder/anaconda) | A Go client library for the Twitter 1.1 API | 1110 | 
| [Kachit/appstore-sdk-go](https://github.com/Kachit/appstore-sdk-go) | Golang SDK for AppStore Connect API (Unofficial) | 2 | 
| [aws/aws-sdk-go](https://github.com/aws/aws-sdk-go) | AWS SDK for the Go programming language. | 7531 | 
| [OTA-Insight/bqwriter](https://github.com/OTA-Insight/bqwriter) | Stream data into Google BigQuery concurrently using InsertAll() or BQ Storage. | 8 | 
| [naegelejd/brewerydb](https://github.com/naegelejd/brewerydb) | Go library for http://www.brewerydb.com/ API | 17 | 
| [andygrunwald/cachet](https://github.com/andygrunwald/cachet) | Go(lang) client library for Cachet (open source status page system). | 90 | 
| [jszwedko/go-circleci](https://github.com/jszwedko/go-circleci) | Go library for interacting with CircleCI | 62 | 
| [Clarifai/clarifai-go](https://github.com/Clarifai/clarifai-go) | DEPRECATED: please use https://github.com/Clarifai/clarifai-go-grpc | 55 | 
| [codeship/codeship-go](https://github.com/codeship/codeship-go) | Go library for accessing the Codeship API v2 | 18 | 
| [coinpaprika/coinpaprika-api-go-client](https://github.com/coinpaprika/coinpaprika-api-go-client) | Go client library for interacting with Coinpaprika's API | 14 | 
| [rinchsan/device-check-go](https://github.com/rinchsan/device-check-go) | :iphone: iOS DeviceCheck SDK for Go - query and modify the per-device bits | 11 | 
| [bwmarrin/discordgo](https://github.com/bwmarrin/discordgo) |  (Golang) Go bindings for Discord | 2859 | 
| [Kachit/dusupay-sdk-go](https://github.com/Kachit/dusupay-sdk-go) | Golang SDK for Dusupay payment gateway API (Unofficial) | 1 | 
| [onrik/ethrpc](https://github.com/onrik/ethrpc) | Golang client for ethereum json rpc api | 224 | 
| [huandu/facebook](https://github.com/huandu/facebook) | A Facebook Graph API SDK For Go. | 1049 | 
| [maddevsio/fcm](https://github.com/maddevsio/fcm) | Firebase Cloud Messaging for application servers implemented using the Go programming language. | 45 | 
| [emiddleton/gads](https://github.com/emiddleton/gads) | Google Adwords API for Go | 49 | 
| [bit4bit/gami](https://github.com/bit4bit/gami) | GO - Asterisk AMI Interface | 31 | 
| [TheOrioli/gcm](https://github.com/TheOrioli/gcm) | Google Cloud Messaging for application servers implemented using the Go programming language. | 30 | 
| [codingsince1985/geo-golang](https://github.com/codingsince1985/geo-golang) | Go library to access geocoding and reverse geocoding APIs | 425 | 
| [google/go-github](https://github.com/google/go-github) | Go library for accessing the GitHub API | 8348 | 
| [shurcooL/githubv4](https://github.com/shurcooL/githubv4) | Package githubv4 is a client library for accessing GitHub GraphQL API v4 (https://docs.github.com/en/graphql). | 864 | 
| [ctreminiom/go-atlassian](https://github.com/ctreminiom/go-atlassian) | ✨ Golang Client Library for Atlassian Cloud. | 42 | 
| [circa10a/go-aws-news](https://github.com/circa10a/go-aws-news) | Go app + library to fetch what's new from AWS | 12 | 
| [axelspringer/go-chronos](https://github.com/axelspringer/go-chronos) | :dancers: Go Chronos 3.x REST API Client | 4 | 
| [PaulRosset/go-hacknews](https://github.com/PaulRosset/go-hacknews) | 📟  Tiny utility Go client for HackerNews API. | 14 | 
| [abdullahselek/go-here](https://github.com/abdullahselek/go-here) | Go client library around the HERE location based APIs. | 10 | 
| [koffeinsource/go-imgur](https://github.com/koffeinsource/go-imgur) | Go library to use the imgur.com API | 21 | 
| [andygrunwald/go-jira](https://github.com/andygrunwald/go-jira) | Go client library for Atlassian Jira | 1091 | 
| [go-lark/lark](https://github.com/go-lark/lark) | An easy-to-use SDK for Feishu and Lark Open Platform (Messaging API only) | 79 | 
| [gambol99/go-marathon](https://github.com/gambol99/go-marathon) | A GO API library for working with Marathon | 195 | 
| [nstratos/go-myanimelist](https://github.com/nstratos/go-myanimelist) | Go library for accessing the MyAnimeList API: https://myanimelist.net/apiconfig/references/api/v2 | 25 | 
| [manuelbcd/go-openproject](https://github.com/manuelbcd/go-openproject) | Go client library for OpenProject | 10 | 
| [rbretecher/go-postman-collection](https://github.com/rbretecher/go-postman-collection) | Go module to work with Postman Collections | 40 | 
| [chriscross0/go-restcountries](https://github.com/chriscross0/go-restcountries) | Go wrapper for the REST Countries API. | 2 | 
| [esurdam/go-sophos](https://github.com/esurdam/go-sophos) | Sophos UTM 9 REST API Client in Golang | 8 | 
| [sergioaugrod/go-sptrans](https://github.com/sergioaugrod/go-sptrans) | Go client library for the SPTrans Olho Vivo API. :bus: | 6 | 
| [esurdam/go-swagger-ui](https://github.com/esurdam/go-swagger-ui) | Golang package which provides http Handlers to serve the swagger ui | 6 | 
| [go-telegraph - Telegraph publishing platform API client.](https://gitlab.com/toby3d/telegraph?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [andygrunwald/go-trending](https://github.com/andygrunwald/go-trending) | Go library for accessing trending repositories and developers at Github. | 118 | 
| [knspriggs/go-twitch](https://github.com/knspriggs/go-twitch) | A golang client for the Twitch v3 API - public APIs only (for now) | 21 | 
| [dghubble/go-twitter](https://github.com/dghubble/go-twitter) | Go Twitter REST and Streaming API v1.1 | 1421 | 
| [hbagdi/go-unsplash](https://github.com/hbagdi/go-unsplash) | Go Client for the Unsplash API  | 58 | 
| [nishanths/go-xkcd](https://github.com/nishanths/go-xkcd) | xkcd.com API client in Go | 43 | 
| [go-yapla - Go client library for the Yapla v2.0 API.](https://git.iglou.eu/Production/go-yapla?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [groovili/gogtrends](https://github.com/groovili/gogtrends) | Unofficial Google Trends API for Go | 58 | 
| [cyruzin/golang-tmdb](https://github.com/cyruzin/golang-tmdb) | This is a Golang wrapper for working with TMDb API. It aims to support version 3. | 50 | 
| [mamal72/golyrics](https://github.com/mamal72/golyrics) | A simple Go package to fetch lyrics from Wikia | 36 | 
| [MonaxGT/gomalshare](https://github.com/MonaxGT/gomalshare) | Go library MalShare API | 9 | 
| [michiwend/gomusicbrainz](https://github.com/michiwend/gomusicbrainz) | a Go (Golang) MusicBrainz WS2 client library - work in progress | 44 | 
| [googleapis/google-api-go-client](https://github.com/googleapis/google-api-go-client) | Auto-generated Google APIs for Go. | 2954 | 
| [chonthu/go-google-analytics](https://github.com/chonthu/go-google-analytics) | Simple Reporting for Google Analytics | 12 | 
| [googleapis/google-cloud-go](https://github.com/googleapis/google-cloud-go) | Google Cloud Client Libraries for Go. | 2795 | 
| [ngs/go-google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) | Go Client Library for G Suite Email Audit API | 7 | 
| [n0madic/google-play-scraper](https://github.com/n0madic/google-play-scraper) | Golang scraper to get data from Google Play Store | 23 | 
| [utekaravinash/gopaapi5](https://github.com/utekaravinash/gopaapi5) | Go Client Library for Amazon's Product Advertising API 5.0 | 11 | 
| [koltyakov/gosip](https://github.com/koltyakov/gosip) | ⚡️ SharePoint authentication, HTTP client & fluent API wrapper for Go (Golang) | 71 | 
| [jsgilmore/gostorm](https://github.com/jsgilmore/gostorm) | GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. | 128 | 
| [andybons/hipchat](https://github.com/andybons/hipchat) | This project implements a Go client library for the Hipchat API. | 104 | 
| [daneharrigan/hipchat](https://github.com/daneharrigan/hipchat) | A golang package to communicate with HipChat over XMPP | 110 | 
| [Henry-Sarabia/igdb](https://github.com/Henry-Sarabia/igdb) | Go client for the Internet Game Database API | 73 | 
| [Icelain/jokeapi](https://github.com/Icelain/jokeapi) | Official golang wrapper for Sv443's jokeapi. | 15 | 
| [Henry-Sarabia/kanka](https://github.com/Henry-Sarabia/kanka) | Go client for the Kanka API | 3 | 
| [chyroc/lark](https://github.com/chyroc/lark) | Feishu/Lark Open API Go SDK, Support ALL Open API and Event Callback. | 135 | 
| [ansd/lastpass-go](https://github.com/ansd/lastpass-go) | Golang client for LastPass | 26 | 
| [clevabit/libgoffi](https://github.com/clevabit/libgoffi) | libgoffi - libffi adapter library for Go | 7 | 
| [Medium/medium-sdk-go](https://github.com/Medium/medium-sdk-go) | A Golang SDK for Medium's OAuth2 API | 131 | 
| [andygrunwald/megos](https://github.com/andygrunwald/megos) | Go(lang) client library for accessing information of an Apache Mesos cluster. | 54 | 
| [minio/minio-go](https://github.com/minio/minio-go) | MinIO Client SDK for Go | 1558 | 
| [dukex/mixpanel](https://github.com/dukex/mixpanel) | Golang Mixpanel Client | 42 | 
| [jellydator/newsapi-go](https://github.com/jellydator/newsapi-go) | Go client for NewsAPI | 0 | 
| [mxpv/patreon-go](https://github.com/mxpv/patreon-go) | Patreon Go API client | 29 | 
| [plutov/paypal](https://github.com/plutov/paypal) | Golang client for PayPal REST API | 473 | 
| [playlyfe/playlyfe-go-sdk](https://github.com/playlyfe/playlyfe-go-sdk) | This is the official Playlyfe Golang Sdk | 1 | 
| [gregdel/pushover](https://github.com/gregdel/pushover) | Go wrapper for the Pushover API | 113 | 
| [dimuska139/rawg-sdk-go](https://github.com/dimuska139/rawg-sdk-go) | This is RAWG SDK GO. This library contains methods for interacting with RAWG API. | 4 | 
| [Omie/rrdaclient](https://github.com/Omie/rrdaclient) | Go bindings for RRDA https://github.com/fcambus/rrda | 8 | 
| [rapito/go-shopify](https://github.com/rapito/go-shopify) | Simple Shopify API for the Go Programming Language | 22 | 
| [rhnvrm/simples3](https://github.com/rhnvrm/simples3) | Simple no frills AWS S3 Golang Library using REST with V4 Signing (without AWS Go SDK) | 85 | 
| [slack-go/slack](https://github.com/slack-go/slack) | Slack API in Go - community-maintained fork created by the original author, @nlopes | 3857 | 
| [sergiotapia/smitego](https://github.com/sergiotapia/smitego) | SmiteGo is an API wrapper for the Smite game from HiRez. It is written in Go! | 10 | 
| [rapito/go-spotify](https://github.com/rapito/go-spotify) | Go library for the Spotify Web API | 42 | 
| [sostronk/go-steam](https://github.com/sostronk/go-steam) | Go library for querying Source servers | 24 | 
| [stripe/stripe-go](https://github.com/stripe/stripe-go) | Go library for the Stripe API.     | 1528 | 
| [farmergreg/textbelt](https://github.com/farmergreg/textbelt) | golang library for textbelt.com | 17 | 
| [nuveo/translate](https://github.com/nuveo/translate) | Go online translation package | 31 | 
| [adlio/trello](https://github.com/adlio/trello) | Trello API wrapper for Go | 194 | 
| [mrbenosborne/tripadvisor-golang](https://github.com/mrbenosborne/tripadvisor-golang) | A TripAdvisor API wrapper for Golang. | 1 | 
| [mattcunningham/gumblr](https://github.com/mattcunningham/gumblr) | A Go Wrapper for the Tumblr v2 API | 6 | 
| [n0madic/twitter-scraper](https://github.com/n0madic/twitter-scraper) | Scrape the Twitter Frontend API without authentication with Golang. | 195 | 
| [bitfield/uptimerobot](https://github.com/bitfield/uptimerobot) | Client library for UptimeRobot v2 API | 46 | 
| [verifid/vl-go](https://github.com/verifid/vl-go) | Go client library around the VerifID identity verification layer API. | 1 | 
| [go-playground/webhooks](https://github.com/go-playground/webhooks) | :fishing_pole_and_fish: Webhook receiver for GitHub, Bitbucket, GitLab, Gogs | 706 | 
| [wit-ai/wit-go](https://github.com/wit-ai/wit-go) | Go client for wit.ai HTTP API | 114 | 
| [brunomvsouza/ynab.go](https://github.com/brunomvsouza/ynab.go) | Go client for the YNAB API. Unofficial. It covers 100% of the resources made available by the YNAB API. | 49 | 
| [gojuno/go-zooz](https://github.com/gojuno/go-zooz) | Zooz API client for Go | 6 | 


# Utilities

| Repository | Description | Stars |
|:---|:---|:---:|
| [topfreegames/apm](https://github.com/topfreegames/apm) | APM is a process manager for Golang applications. | 154 | 
| [icza/backscanner](https://github.com/icza/backscanner) | A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. | 35 | 
| [wesovilabs/beyond](https://github.com/wesovilabs/beyond) | The Go library that will drive you to AOP world! | 46 | 
| [Henry-Sarabia/blank](https://github.com/Henry-Sarabia/blank) | Detect blank strings or remove whitespace from strings | 7 | 
| [sinhashubham95/bleep](https://github.com/sinhashubham95/bleep) | OS Signal Handlers in Go | 4 | 
| [tmrts/boilr](https://github.com/tmrts/boilr) | :zap: boilerplate template manager that generates files or directories from template repositories | 1443 | 
| [miniscruff/changie](https://github.com/miniscruff/changie) | Automated changelog tool for preparing releases with lots of customization options | 168 | 
| [antham/chyle](https://github.com/antham/chyle) | Changelog generator : use a git repository and various data sources and publish the result on external services | 144 | 
| [cep21/circuit](https://github.com/cep21/circuit) | An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. | 635 | 
| [rubyist/circuitbreaker](https://github.com/rubyist/circuitbreaker) | Circuit Breakers in Go | 985 | 
| [golang-design/clipboard](https://github.com/golang-design/clipboard) | 📋 cross-platform clipboard package that supports accessing text and image in Go (macOS/Linux/Windows/Android/iOS)  | 176 | 
| [jonboulle/clockwork](https://github.com/jonboulle/clockwork) | a fake clock for golang | 376 | 
| [commander-cli/cmd](https://github.com/commander-cli/cmd) | A simple package to execute shell commands on linux, windows and osx | 87 | 
| [txgruppi/command](https://github.com/txgruppi/command) | Command pattern for Go with thread safe serial and parallel dispatcher | 13 | 
| [gotidy/copy](https://github.com/gotidy/copy) | Package for fast copying structs of different types | 15 | 
| [jutkko/copy-pasta](https://github.com/jutkko/copy-pasta) | Universal copy paste service, works across different machines! | 50 | 
| [biter777/countries](https://github.com/biter777/countries) | Countries - ISO 3166 (ISO3166-1, ISO3166, Digit, Alpha-2 and Alpha-3) countries codes and names (on eng and rus), ISO 4217 currency designators, ITU-T E.164 IDD calling phone codes, countries capitals, UN M.49 regions codes, ccTLD countries domains, IOC/NOC and FIFA letters codes, VERY FAST, NO maps[], NO slices[], NO init() funcs, NO external links/files/data, NO interface{}, NO specific dependencies, Databases/JSON/GOB/XML/CSV compatible, Emoji countries flags and currencies support, full support ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and ccTLD standarts. | 152 | 
| [create-go-app/cli](https://github.com/create-go-app/cli) | ✨ Create a new production-ready project with backend, frontend and deploy automation by running one CLI command! | 1325 | 
| [Gituser143/cryptgo](https://github.com/Gituser143/cryptgo) | A terminal application to watch crypto prices! | 105 | 
| [bcicen/ctop](https://github.com/bcicen/ctop) | Top-like interface for container metrics | 12482 | 
| [posener/ctxutil](https://github.com/posener/ctxutil) | utils for Go context | 18 | 
| [shockerli/cvt](https://github.com/shockerli/cvt) | Easy and safe convert any value to another type. Go 任意数据类型安全转换 | 15 | 
| [nikogura/dbt](https://github.com/nikogura/dbt) | Dynamic Binary Toolkit- A framework for running self-updating signed binaries from a central, trusted repository. | 46 | 
| [vrecan/death](https://github.com/vrecan/death) | Managing go application shutdown with signals. | 178 | 
| [ulule/deepcopier](https://github.com/ulule/deepcopier) | simple struct copying for golang | 380 | 
| [derekparker/delve](https://github.com/derekparker/delve) | Delve is a debugger for the Go programming language. | 471 | 
| [kirillDanshin/dlog](https://github.com/kirillDanshin/dlog) | Simple build-time controlled debug log with ability to log where the logger was called | 15 | 
| [reugn/equalizer](https://github.com/reugn/equalizer) | A rate limiters package for Go | 35 | 
| [cristianoliveira/ergo](https://github.com/cristianoliveira/ergo) | The management of multiple apps running over different ports made easy | 506 | 
| [nullne/evaluator](https://github.com/nullne/evaluator) |  | 33 | 
| [h2non/filetype](https://github.com/h2non/filetype) | Fast, dependency-free Go package to infer binary file types based on the magic numbers header signature | 1504 | 
| [yaronsumel/filler](https://github.com/yaronsumel/filler) | fill struct data easily with fill tags | 16 | 
| [gookit/filter](https://github.com/gookit/filter) | ⏳ Provide filtering, sanitizing, and conversion of Golang data. 提供对Golang数据的过滤，净化，转换。 | 56 | 
| [junegunn/fzf](https://github.com/junegunn/fzf) | :cherry_blossom: A command-line fuzzy finder | 43032 | 
| [go-playground/generate](https://github.com/go-playground/generate) | :runner:runs go generate recursively on a specified path or environment variable and can filter by regex | 25 | 
| [antham/ghokin](https://github.com/antham/ghokin) | Parallelized formatter with no external dependencies for gherkin (cucumber, behat...) | 23 | 
| [git-time-metric/gtm](https://github.com/git-time-metric/gtm) | Simple, seamless, lightweight time tracking for Git | 894 | 
| [sinhashubham95/go-actuator](https://github.com/sinhashubham95/go-actuator) | Golang production-ready features | 5 | 
| [asticode/go-astitodo](https://github.com/asticode/go-astitodo) | Parse TODOs in your GO code | 58 | 
| [wendigo/go-bind-plugin](https://github.com/wendigo/go-bind-plugin) | go-bind-plugin generates API for exported plugin symbols (-buildmode=plugin) - go1.8+ only (http://golang.org/pkg/plugin) | 178 | 
| [gabstv/go-bsdiff](https://github.com/gabstv/go-bsdiff) | Pure Go bsdiff and bspatch libraries and CLI tools. | 123 | 
| [prashantgupta24/go-clip](https://github.com/prashantgupta24/go-clip) | A minimalistic clipboard manager for Mac. | 8 | 
| [Eun/go-convert](https://github.com/Eun/go-convert) | Convert a value into another type | 15 | 
| [mikekonan/go-countries](https://github.com/mikekonan/go-countries) |  | 9 | 
| [ungerik/go-dry](https://github.com/ungerik/go-dry) | DRY (don't repeat yourself) package for Go | 479 | 
| [thoas/go-funk](https://github.com/thoas/go-funk) | A modern Go utility library which provides helpers (map, find, contains, filter, ...) | 3380 | 
| [Talento90/go-health](https://github.com/Talento90/go-health) | :heart: Health check your applications and dependencies | 87 | 
| [mozillazg/go-httpheader](https://github.com/mozillazg/go-httpheader) | A Go library for encoding structs into Header fields. | 38 | 
| [viney-shih/go-lock](https://github.com/viney-shih/go-lock) | go-lock is a lock library implementing read-write mutex and read-write trylock without starvation | 61 | 
| [chenquan/go-pkg](https://github.com/chenquan/go-pkg) | A go toolkit. | 5 | 
| [mvmaasakkers/go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) | Problem json implementation (https://tools.ietf.org/html/rfc7807) package for go | 12 | 
| [beefsack/go-rate](https://github.com/beefsack/go-rate) | A timed rate limiter for Go | 353 | 
| [kenkyu392/go-safe](https://github.com/kenkyu392/go-safe) | This Go package provides a sandbox for the safe execution of panic-inducing programs | 4 | 
| [ikeikeikeike/go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) | go-sitemap-generator is the easiest way to generate Sitemaps in Go | 172 | 
| [sadlil/go-trigger](https://github.com/sadlil/go-trigger) | A Global event triggerer for golang. Defines functions as event with id string. Trigger the event anywhere from your project. | 225 | 
| [mikekonan/go-types](https://github.com/mikekonan/go-types) | Library providing opanapi3 and Go types for store/validation and transfer of ISO-4217, ISO-3166, and other types. | 12 | 
| [carlescere/goback](https://github.com/carlescere/goback) | Golang simple exponential backoff package. | 44 | 
| [zerosnake0/goctx](https://github.com/zerosnake0/goctx) | Get your context value faster | 2 | 
| [VividCortex/godaemon](https://github.com/VividCortex/godaemon) | Daemonize Go applications deviously. | 487 | 
| [dropbox/godropbox](https://github.com/dropbox/godropbox) | Common libraries for writing Go services/applications. | 4012 | 
| [cosiner/gohper](https://github.com/cosiner/gohper) | [UNMATAINED] common libs here. | 256 | 
| [msempere/golarm](https://github.com/msempere/golarm) | Fire alarms with system events | 45 | 
| [mlimaloureiro/golog](https://github.com/mlimaloureiro/golog) | Easy and simple CLI time tracker for your tasks | 55 | 
| [bndr/gopencils](https://github.com/bndr/gopencils) | Easily consume REST APIs with Go (golang) | 437 | 
| [michiwend/goplaceholder](https://github.com/michiwend/goplaceholder) | a small golang lib to generate placeholder images | 24 | 
| [philipjkim/goreadability](https://github.com/philipjkim/goreadability) | Webpage summary extractor using Facebook Open Graph and arc90's readability | 63 | 
| [goreleaser/goreleaser](https://github.com/goreleaser/goreleaser) | Deliver Go binaries as fast and easily as possible | 9807 | 
| [qax-os/goreporter](https://github.com/qax-os/goreporter) | A Golang tool that does static analysis, unit testing, code review and generate code quality report. | 2972 | 
| [linxGnu/goseaweedfs](https://github.com/linxGnu/goseaweedfs) | A complete Golang client for SeaweedFS | 93 | 
| [ik5/gostrutils](https://github.com/ik5/gostrutils) | Collections of string utils I have created over the years | 34 | 
| [subosito/gotenv](https://github.com/subosito/gotenv) | Load environment variables from `.env` or `io.Reader` in Go. | 214 | 
| [maja42/goval](https://github.com/maja42/goval) | Expression evaluation in golang | 70 | 
| [tenntenn/gpath](https://github.com/tenntenn/gpath) | gpath is a Go package to access a field by a path using reflect pacakge | 39 | 
| [pesos/grofer](https://github.com/pesos/grofer) | A system and resource monitoring tool written in Golang! | 191 | 
| [novalagung/gubrak](https://github.com/novalagung/gubrak) | ⚙️ Golang functional utility library with syntactic sugar. It's like lodash, but for Go | 394 | 
| [miguelpragier/handy](https://github.com/miguelpragier/handy) | GO Golang Utilities and helpers like validators and string formatters | 66 | 
| [guumaster/hostctl](https://github.com/guumaster/hostctl) | Your dev tool to manage /etc/hosts like a pro! | 719 | 
| [htcat/htcat](https://github.com/htcat/htcat) | Parallel and Pipelined HTTP GET Utility | 549 | 
| [github/hub](https://github.com/github/hub) | A command-line tool that makes git easier to use with GitHub. | 21643 | 
| [afex/hystrix-go](https://github.com/afex/hystrix-go) | Netflix's Hystrix latency and fault tolerance library, for Go  | 3574 | 
| [immortal/immortal](https://github.com/immortal/immortal) | ⭕  A *nix cross-platform (OS agnostic) supervisor | 727 | 
| [mengzhuo/intrinsic](https://github.com/mengzhuo/intrinsic) | Provide Golang native SIMD intrinsics on x86/amd64 platform | 44 | 
| [clevergo/jsend](https://github.com/clevergo/jsend) | :100: JSend's implementation writen in Go(golang) | 14 | 
| [gsamokovarov/jump](https://github.com/gsamokovarov/jump) | Jump helps you navigate faster by learning your habits. ✌️ | 1325 | 
| [wesovilabs/koazee](https://github.com/wesovilabs/koazee) | A StreamLike, Immutable, Lazy Loading and smart Golang Library to deal with slices. | 490 | 
| [aplescia/lets-go](https://github.com/aplescia/lets-go) | Go module that provides common utilities for Cloud Native development | 3 | 
| [mennanov/limiters](https://github.com/mennanov/limiters) | Golang rate limiters for distributed applications | 86 | 
| [samber/lo](https://github.com/samber/lo) | 💥  A Lodash-style Go library based on Go 1.18+ Generics (map, filter, contains, find...) | 4164 | 
| [jaschaephraim/lrserver](https://github.com/jaschaephraim/lrserver) | LiveReload server for Go [golang] | 120 | 
| [alajmo/mani](https://github.com/alajmo/mani) | CLI tool to help you manage multiple repositories | 186 | 
| [minio/mc](https://github.com/minio/mc) | MinIO Client is a replacement for ls, cp, mkdir, diff and rsync commands for filesystems and object storage. | 2070 | 
| [imdario/mergo](https://github.com/imdario/mergo) | Mergo: merging Go structs and maps since 2013. | 1895 | 
| [zRedShift/mimemagic](https://github.com/zRedShift/mimemagic) | Powerful and versatile MIME sniffing package using pre-compiled glob patterns, magic number signatures, XML document namespaces, and tree magic for mounted volumes, generated from the XDG shared-mime-info database. | 74 | 
| [aofei/mimesniffer](https://github.com/aofei/mimesniffer) | A MIME type sniffer for Go. | 20 | 
| [gabriel-vasile/mimetype](https://github.com/gabriel-vasile/mimetype) | A fast Golang library for media type and file extension detection, based on magic numbers | 694 | 
| [tdewolff/minify](https://github.com/tdewolff/minify) | Go minifiers for web formats | 2914 | 
| [icza/minquery](https://github.com/icza/minquery) | MongoDB / mgo query that supports efficient pagination (cursors to continue listing documents where we left off). | 59 | 
| [StabbyCutyou/moldova](https://github.com/StabbyCutyou/moldova) | A lightweight templating system for generating random data | 160 | 
| [davrodpin/mole](https://github.com/davrodpin/mole) | CLI application to create ssh tunnels focused on resiliency and user experience. | 1548 | 
| [gobeam/mongo-go-pagination](https://github.com/gobeam/mongo-go-pagination) | Golang Mongodb Pagination for official mongodb/mongo-go-driver package which supports both normal queries and Aggregation pipelines with all information like Total records, Page, Per Page, Previous, Next, Total Page and query results. | 94 | 
| [linxGnu/mssqlx](https://github.com/linxGnu/mssqlx) | Database client library, proxy for any master slave, master master structures. Lightweight, performant and auto balancing in mind. | 94 | 
| [VividCortex/multitick](https://github.com/VividCortex/multitick) | A multiplexor for aligned time.Time tickers in Go | 64 | 
| [inancgumus/myhttp](https://github.com/inancgumus/myhttp) | Simplest HTTP GET requester for Go with timeout support | 35 | 
| [e-dard/netbug](https://github.com/e-dard/netbug) | Package netbug provides a handler for registering profilers on your own ServeMux. | 69 | 
| [chrispassas/nfdump](https://github.com/chrispassas/nfdump) | NFDump File Reader | 6 | 
| [pokanop/nostromo](https://github.com/pokanop/nostromo) | CLI for building powerful aliases | 110 | 
| [rekby/objwalker](https://github.com/rekby/objwalker) |  | 1 | 
| [xta/okrun](https://github.com/xta/okrun) | ok, run your gofile | 15 | 
| [btnguyen2k/olaf](https://github.com/btnguyen2k/olaf) | Twitter Snowflake implemented in Go | 3 | 
| [adelowo/onecache](https://github.com/adelowo/onecache) | One caching API, Multiple backends | 127 | 
| [maruel/panicparse](https://github.com/maruel/panicparse) | Crash your app in style (Golang) | 3044 | 
| [alexpantyukhin/go-pattern-match](https://github.com/alexpantyukhin/go-pattern-match) | Pattern matchings for Go. | 174 | 
| [peco/peco](https://github.com/peco/peco) | Simplistic interactive filtering tool | 6825 | 
| [arthurkushman/pgo](https://github.com/arthurkushman/pgo) | Go library for PHP community with convenient functions | 64 | 
| [VividCortex/pm](https://github.com/VividCortex/pm) | Processlist manager with TCP listener | 76 | 
| [gotidy/ptr](https://github.com/gotidy/ptr) | Contains functions for simplified creation of pointers from constants of basic types | 13 | 
| [zpatrick/rclient](https://github.com/zpatrick/rclient) | Minimalistic REST client for Go applications | 32 | 
| [ssgreg/repeat](https://github.com/ssgreg/repeat) | Go implementation of different backoff strategies useful for retrying operations and heartbeating. | 78 | 
| [mozillazg/request](https://github.com/mozillazg/request) | A developer-friendly HTTP request library for Gopher. | 407 | 
| [abo/rerate](https://github.com/abo/rerate) | redis-based rate counter and rate limiter | 22 | 
| [ivpusic/rerun](https://github.com/ivpusic/rerun) | Configurable recompiling and rerunning go apps when source changes | 161 | 
| [edermanoel94/rest-go](https://github.com/edermanoel94/rest-go) | A package that provide many helpful methods for working with rest api. | 16 | 
| [kamilsk/retry](https://github.com/kamilsk/retry) | ♻️ The most advanced interruptible mechanism to perform actions repetitively until successful. | 318 | 
| [percolate/retry](https://github.com/percolate/retry) | Percolate's Go retry package | 8 | 
| [thedevsaddam/retry](https://github.com/thedevsaddam/retry) | Simple and easy retry mechanism package for Go | 54 | 
| [shafreeck/retry](https://github.com/shafreeck/retry) | A pretty simple library to ensure your work to be done | 10 | 
| [rafaeljesus/retry-go](https://github.com/rafaeljesus/retry-go) | Retrying made simple and easy for golang :repeat:  | 43 | 
| [VividCortex/robustly](https://github.com/VividCortex/robustly) | Run functions resiliently in Go, catching and restarting panics | 151 | 
| [ferama/rospo](https://github.com/ferama/rospo) | 🐸 Simple, reliable, persistent ssh tunnels with embedded ssh server | 159 | 
| [blockloop/scan](https://github.com/blockloop/scan) | Scan database/sql rows directly to structs, slices, and primitive types | 251 | 
| [georgysavva/scany](https://github.com/georgysavva/scany) | Library for scanning data from a database into Go structs and more | 519 | 
| [syntaqx/serve](https://github.com/syntaqx/serve) | 🍽️ a static http server anywhere you need one. | 260 | 
| [nofeaturesonlybugs/set](https://github.com/nofeaturesonlybugs/set) | Package set is a small wrapper around the official reflect package that facilitates loose type conversion and assignment into native Go types. | 29 | 
| [ztrue/shutdown](https://github.com/ztrue/shutdown) | Golang app shutdown hooks. | 29 | 
| [chrispassas/silk](https://github.com/chrispassas/silk) | Read Silk Flow Files | 11 | 
| [psampaz/slice](https://github.com/psampaz/slice) | Type-safe functions for common Go slice operations | 47 | 
| [Henry-Sarabia/sliceconv](https://github.com/Henry-Sarabia/sliceconv) | Slice conversion between primitive types | 8 | 
| [leaanthony/slicer](https://github.com/leaanthony/slicer) | Utility class for handling slices | 32 | 
| [jfcg/sorty](https://github.com/jfcg/sorty) | :zap: Fast Concurrent / Parallel Sorting in Go | 98 | 
| [jmoiron/sqlx](https://github.com/jmoiron/sqlx) | general purpose extensions to golang's database/sql | 11631 | 
| [shoobyban/sshman](https://github.com/shoobyban/sshman) | SSH Manager - manage authorized_keys file on remote servers | 29 | 
| [janiltonmaciel/statiks](https://github.com/janiltonmaciel/statiks) | Fast, zero-configuration, static HTTP filer server. | 9 | 
| [asdine/storm](https://github.com/asdine/storm) | Simple and powerful toolkit for BoltDB | 1860 | 
| [PumpkinSeed/structs](https://github.com/PumpkinSeed/structs) | Golang struct operations. | 20 | 
| [yudppp/throttle](https://github.com/yudppp/throttle) | lodash throttle like Go library | 28 | 
| [andy2046/tik](https://github.com/andy2046/tik) | hierarchical timing wheel | 3 | 
| [cyruzin/tome](https://github.com/cyruzin/tome) | Package tome was designed to paginate simple RESTful APIs. | 29 | 
| [viant/toolbox](https://github.com/viant/toolbox) | Toolbox - go utility library | 174 | 
| [alxrm/ugo](https://github.com/alxrm/ugo) | Simple and expressive toolbox written in Go | 26 | 
| [esemplastic/unis](https://github.com/esemplastic/unis) | UNIS: A Common Architecture for String Utilities within the Go Programming Language. | 69 | 
| [xo/usql](https://github.com/xo/usql) | Universal command-line interface for SQL databases | 6981 | 
| [shomali11/util](https://github.com/shomali11/util) | A collection of useful utility functions | 244 | 
| [asciimoo/wuzz](https://github.com/asciimoo/wuzz) | Interactive cli tool for HTTP inspection | 9931 | 
| [monmohan/xferspdy](https://github.com/monmohan/xferspdy) | Xferspdy provides binary diff and patch library in golang. [Mentioned in Awesome Go, https://github.com/avelino/awesome-go] | 91 | 


# UUID

| Repository | Description | Stars |
|:---|:---|:---:|
| [Hart87/goflake](https://github.com/Hart87/goflake) | A highly scalable and serverless unique ID generator for use in distributed systems. Written in GoLang. Inspired by Twitters Snowflake. | 10 | 
| [JakeHL/Goid](https://github.com/JakeHL/Goid) | A UUIDv4 generation package written in go | 31 | 
| [twharmon/gouid](https://github.com/twharmon/gouid) | Fast, dependable universally unique ids | 13 | 
| [aidarkhanov/nanoid](https://github.com/aidarkhanov/nanoid) | A tiny and fast Go unique string generator | 45 | 
| [muyo/sno](https://github.com/muyo/sno) | Compact, sortable and fast unique IDs with embedded metadata. | 62 | 
| [oklog/ulid](https://github.com/oklog/ulid) | Universally Unique Lexicographically Sortable Identifier (ULID) in Go | 2615 | 
| [uniq - No hassle safe, fast unique identifiers with commands.](https://gitlab.com/skilstak/code/go/uniq?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [agext/uuid](https://github.com/agext/uuid) | Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. | 14 | 
| [gofrs/uuid](https://github.com/gofrs/uuid) | A UUID package originally forked from github.com/satori/go.uuid | 1096 | 
| [google/uuid](https://github.com/google/uuid) | Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services. | 3451 | 
| [edwingeng/wuid](https://github.com/edwingeng/wuid) | An extremely fast UUID alternative written in golang | 457 | 
| [rs/xid](https://github.com/rs/xid) | xid is a globally unique id generator thought for the web | 2471 | 


# Validation

| Repository | Description | Stars |
|:---|:---|:---:|
| [osamingo/checkdigit](https://github.com/osamingo/checkdigit) | Provide check digit algorithms and calculators written in Go | 88 | 
| [guiferpa/gody](https://github.com/guiferpa/gody) | :balloon: A lightweight struct validator for Go | 55 | 
| [twharmon/govalid](https://github.com/twharmon/govalid) | Struct validation using tags | 24 | 
| [asaskevich/govalidator](https://github.com/asaskevich/govalidator) | [Go] Package of validators and sanitizers for strings, numerics, slices and structs | 5284 | 
| [thedevsaddam/govalidator](https://github.com/thedevsaddam/govalidator) | Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. | 1071 | 
| [faceair/jio](https://github.com/faceair/jio) | jio is a json schema validator similar to joi | 64 | 
| [go-ozzo/ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | An idiomatic Go (golang) validation package. Supports configurable and extensible validation rules (validators) using normal language constructs instead of error-prone struct tags. | 2602 | 
| [thazelart/terraform-validator](https://github.com/thazelart/terraform-validator) | A norms and conventions validator for Terraform | 76 | 
| [gookit/validate](https://github.com/gookit/validate) | ⚔ Go package for data validation and filtering. support Map, Struct, Form data. Go通用的数据验证与过滤库，使用简单，内置大部分常用验证、过滤器，支持自定义验证器、自定义消息、字段翻译。 | 529 | 
| [gobuffalo/validate](https://github.com/gobuffalo/validate) | This package provides a framework for writing validations for Go applications. | 66 | 
| [go-playground/validator](https://github.com/go-playground/validator) | :100:Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving | 9958 | 
| [go-the-way/validator](https://github.com/go-the-way/validator) | A lightweight model validator written in Go. | 2 | 


# Version Control

| Repository | Description | Stars |
|:---|:---|:---:|
| [jfrog/froggit-go](https://github.com/jfrog/froggit-go) | Froggit-Go is a universal Go library, allowing to perform actions on VCS providers. | 13 | 
| [rjeczalik/gh](https://github.com/rjeczalik/gh) | Scriptable server and net/http middleware for GitHub Webhooks. | 76 | 
| [libgit2/git2go](https://github.com/libgit2/git2go) | Git to Go; bindings for libgit2. Like McDonald's but tastier. | 1708 | 
| [gabyx/Githooks](https://github.com/gabyx/Githooks) | 🦎 Githooks: per-repo and shared Git hooks with version control and auto update.  | 37 | 
| [profclems/glab](https://github.com/profclems/glab) | A GitLab CLI tool bringing GitLab to your command line | 1802 | 
| [go-git/go-git](https://github.com/go-git/go-git) | A highly extensible Git implementation in pure Go. | 3224 | 
| [sourcegraph/go-vcs](https://github.com/sourcegraph/go-vcs) | manipulate and inspect VCS repositories in Go | 74 | 
| [src-d/hercules](https://github.com/src-d/hercules) | Gaining advanced insights from Git repository history. | 1705 | 
| [beyang/hgo](https://github.com/beyang/hgo) | Hgo is a collection of Go packages providing read-access to local Mercurial repositories. | 13 | 


# Video

| Repository | Description | Stars |
|:---|:---|:---:|
| [3d0c/gmf](https://github.com/3d0c/gmf) | Go Media Framework | 746 | 
| [asticode/go-astisub](https://github.com/asticode/go-astisub) | Manipulate subtitles in GO (.srt, .ssa/.ass, .stl, .ttml, .vtt (webvtt), teletext, etc.) | 380 | 
| [asticode/go-astits](https://github.com/asticode/go-astits) | Demux and mux MPEG Transport Streams (.ts) natively in GO | 407 | 
| [quangngotan95/go-m3u8](https://github.com/quangngotan95/go-m3u8) | Parse and generate m3u8 playlists for Apple HTTP Live Streaming (HLS) in Golang (ported from gem https://github.com/sethdeckard/m3u8) | 89 | 
| [unki2aut/go-mpd](https://github.com/unki2aut/go-mpd) | Go library for parsing and generating MPEG-DASH Media Presentation Description (MPD) files | 11 | 
| [giorgisio/goav](https://github.com/giorgisio/goav) | Golang bindings for FFmpeg | 1834 | 
| [aler9/gortsplib](https://github.com/aler9/gortsplib) | RTSP 1.0 client and server library for the Go programming language | 213 | 
| [ziutek/gst](https://github.com/ziutek/gst) | Go bindings for GStreamer (retired: currently I don't use/develop this package) | 166 | 
| [wargarblgarbl/libgosubs](https://github.com/wargarblgarbl/libgosubs) | golang library to read and write various subtitle formats | 18 | 
| [adrg/libvlc-go](https://github.com/adrg/libvlc-go) | Go bindings for libVLC and high-level media player interface | 287 | 
| [grafov/m3u8](https://github.com/grafov/m3u8) | Parser and generator of M3U8-playlists for Apple HLS. Library for Go language. :cinema: | 908 | 
| [korandiz/v4l](https://github.com/korandiz/v4l) | Facade to the Video4Linux video capture interface.  | 65 | 


# Web Frameworks

| Repository | Description | Stars |
|:---|:---|:---:|
| [aah - Scalable, performant, rapid development Web framework for Go.](https://aahframework.org?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [aerogo/aero](https://github.com/aerogo/aero) | :bullettrain_side: High-performance web server for Go. | 445 | 
| [aofei/air](https://github.com/aofei/air) | An ideally refined web framework for Go. | 413 | 
| [appist/appy](https://github.com/appist/appy) | An opinionated productive web framework that helps scaling business easier. | 116 | 
| [n4Zz2/banjo](https://github.com/n4Zz2/banjo) | BANjO is a simple web framework written in Go (golang) | 19 | 
| [beego/beego](https://github.com/beego/beego) | beego is an open-source, high-performance web framework for the Go programming language. | 27914 | 
| [Buffalo - Bringing the productivity of Rails to Go!](https://gobuffalo.io?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [Confetti Framework - Confetti is a Go web application framework with an expressive, elegant syntax. Confetti combines the elegance of Laravel and the simplicity of Go.](https://confetti-framework.github.io/docs/?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [labstack/echo](https://github.com/labstack/echo) | High performance, minimalist Go web framework | 21932 | 
| [gofiber/fiber](https://github.com/gofiber/fiber) | ⚡️ Express inspired web framework written in Go | 19146 | 
| [zpatrick/fireball](https://github.com/zpatrick/fireball) | Go web framework with a natural feel | 57 | 
| [i-love-flamingo/flamingo](https://github.com/i-love-flamingo/flamingo) | Flamingo Framework and Core Library. Flamingo is a go based framework for pluggable web projects. It is used to build scalable and maintainable (web)applications. | 285 | 
| [i-love-flamingo/flamingo-commerce](https://github.com/i-love-flamingo/flamingo-commerce) | Flexible E-Commerce Framework on top of Flamingo. Used to build E-Commerce "Portals" and connect it with the help of individual Adapters to other services.  | 301 | 
| [gogearbox/gearbox](https://github.com/gogearbox/gearbox) | Gearbox :gear: is a web framework written in Go with a focus on high performance | 615 | 
| [gin-gonic/gin](https://github.com/gin-gonic/gin) | Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin. | 56836 | 
| [xxjwxc/ginrpc](https://github.com/xxjwxc/ginrpc) | gin auto binding,grpc, and annotated route,gin 注解路由, grpc,自动参数绑定工具 | 210 | 
| [nytimes/gizmo](https://github.com/nytimes/gizmo) | A Microservice Toolkit from The New York Times | 3596 | 
| [ant0ine/go-json-rest](https://github.com/ant0ine/go-json-rest) | A quick and easy way to setup a RESTful JSON API | 3495 | 
| [ungerik/go-rest](https://github.com/ungerik/go-rest) | A small and evil REST framework for Go | 125 | 
| [goadesign/goa](https://github.com/goadesign/goa) | Design-based APIs and microservices in Go | 4608 | 
| [goa-go/goa](https://github.com/goa-go/goa) | Goa is a  web framework based on middleware, like koa.js. | 46 | 
| [fulldump/golax](https://github.com/fulldump/golax) | Golax, a go implementation for the Lax framework. | 74 | 
| [dinever/golf](https://github.com/dinever/golf) | :golf: The Golf web framework | 253 | 
| [rainycape/gondola](https://github.com/rainycape/gondola) | The web framework for writing faster sites, faster | 309 | 
| [mustafaakin/gongular](https://github.com/mustafaakin/gongular) | A different approach to Go web frameworks | 448 | 
| [gotuna/gotuna](https://github.com/gotuna/gotuna) | GoTuna a lightweight web framework for Go with mux router, middlewares, user sessions, templates, embedded views, and static file server. | 39 | 
| [twharmon/goweb](https://github.com/twharmon/goweb) | Lightweight web framework based on net/http. | 28 | 
| [go-goyave/goyave](https://github.com/go-goyave/goyave) | 🍐 Elegant Golang REST API Framework | 954 | 
| [hidevopsio/hiboot](https://github.com/hidevopsio/hiboot) | hiboot is a high performance web and cli application framework with dependency injection support | 164 | 
| [danielgtaylor/huma](https://github.com/danielgtaylor/huma) | Huma REST/GraphQL API Framework for Golang with OpenAPI 3 | 60 | 
| [go-macaron/macaron](https://github.com/go-macaron/macaron) | Package macaron is a high productive and modular web framework in Go. | 3265 | 
| [paulbellamy/mango](https://github.com/paulbellamy/mango) | Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. | 362 | 
| [claygod/microservice](https://github.com/claygod/microservice) | This library provides a simple microservice framework based on clean architecture principles with a working example implemented. | 91 | 
| [ivpusic/neo](https://github.com/ivpusic/neo) | Go Web Framework | 416 | 
| [beatlabs/patron](https://github.com/beatlabs/patron) | Microservice framework following best cloud practices with a focus on productivity. | 94 | 
| [resoursea/api](https://github.com/resoursea/api) | A REST framework for quickly writing resource based services in Golang. | 32 | 
| [REST Layer - Framework to build REST/GraphQL API on top of databases with mostly configuration over code.](https://rest-layer.io?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [revel/revel](https://github.com/revel/revel) | A high productivity, full-stack web framework for the Go language. | 12518 | 
| [goanywhere/rex](https://github.com/goanywhere/rex) | Pleasures for Web in Golang | 32 | 
| [rookie-ninja/rk-boot](https://github.com/rookie-ninja/rk-boot) | Bootstrapper for golang application. See https://rkdev.info/docs/ for details. | 178 | 
| [gookit/rux](https://github.com/gookit/rux) | ⚡ Rux is an simple and fast web framework. support route group, param route binding, middleware, compatible http.Handler interface. 简单且快速的 Go api/web 框架，支持路由分组，路由参数绑定，中间件，兼容 http.Handler 接口 | 73 | 
| [lunny/tango](https://github.com/lunny/tango) | This is only a mirror and Moved to https://gitea.com/lunny/tango | 835 | 
| [rcrowley/go-tigertonic](https://github.com/rcrowley/go-tigertonic) | A Go framework for building JSON web services inspired by Dropwizard | 999 | 
| [uadmin/uadmin](https://github.com/uadmin/uadmin) | The web framework for Golang | 182 | 
| [gernest/utron](https://github.com/gernest/utron) | A lightweight MVC framework for Go(Golang) | 2215 | 
| [aisk/vox](https://github.com/aisk/vox) | Simple and lightweight Go web framework inspired by koa | 76 | 
| [bnkamalesh/webgo](https://github.com/bnkamalesh/webgo) | A microframework to build web apps; with handler chaining, middleware support, and most of all; standard library compliant HTTP handlers(i.e. http.HandlerFunc). | 230 | 
| [yarf-framework/yarf](https://github.com/yarf-framework/yarf) | Yet Another REST Framework | 65 | 


# Actual middlewares

| Repository | Description | Stars |
|:---|:---|:---:|


# Libraries for creating HTTP middlewares

| Repository | Description | Stars |
|:---|:---|:---:|


# Routers

| Repository | Description | Stars |
|:---|:---|:---:|


# WebAssembly

| Repository | Description | Stars |
|:---|:---|:---:|


# Windows

| Repository | Description | Stars |
|:---|:---|:---:|


# XML

| Repository | Description | Stars |
|:---|:---|:---:|
| [XML-Comp/XML-Comp](https://github.com/XML-Comp/XML-Comp) | Compare ANY markup documents. | 16 | 
| [sbabiv/xml2map](https://github.com/sbabiv/xml2map) | XML to MAP converter written Golang | 37 | 
| [shabbyrobe/xmlwriter](https://github.com/shabbyrobe/xmlwriter) | xmlwriter is a pure-Go library providing procedural XML generation based on libxml2's xmlwriter module | 21 | 
| [antchfx/xpath](https://github.com/antchfx/xpath) | XPath package for Golang, supports HTML, XML, JSON document query. | 463 | 
| [antchfx/xquery](https://github.com/antchfx/xquery) | Extract data or evaluate value from HTML/XML documents using XPath | 156 | 
| [miku/zek](https://github.com/miku/zek) | Generate a Go struct from XML. | 543 | 


# Zero Trust

| Repository | Description | Stars |
|:---|:---|:---:|
| [sigstore/cosign](https://github.com/sigstore/cosign) | Container Signing | 1779 | 
| [in-toto/in-toto-golang](https://github.com/in-toto/in-toto-golang) | A Go implementation of in-toto. in-toto is a framework to protect software supply chain integrity. | 51 | 
| [philips-labs/spiffe-vault](https://github.com/philips-labs/spiffe-vault) | Integrates Spiffe and Vault to have secretless authentication | 16 | 
| [spiffe/spire](https://github.com/spiffe/spire) | The SPIFFE Runtime Environment | 1071 | 


# Code Analysis

| Repository | Description | Stars |
|:---|:---|:---:|
| [bradleyfalzon/apicompat](https://github.com/bradleyfalzon/apicompat) | apicompat checks recent changes to a Go project for backwards incompatible changes | 176 | 
| [Checkmarx/chainjacking](https://github.com/Checkmarx/chainjacking) | Find which of your go lang direct GitHub dependencies is susceptible to ChainJacking attack | 18 | 
| [mibk/dupl](https://github.com/mibk/dupl) | a tool for code clone detection | 277 | 
| [kisielk/errcheck](https://github.com/kisielk/errcheck) | errcheck checks that you checked errors. | 1807 | 
| [davecheney/gcvis](https://github.com/davecheney/gcvis) | Visualise Go program GC trace data in real time | 1055 | 
| [qiniu/checkstyle](https://github.com/qiniu/checkstyle) | checkstyle for go | 119 | 
| [roblaszczak/go-cleanarch](https://github.com/roblaszczak/go-cleanarch) | Clean architecture validator for go, like a The Dependency Rule and interaction between packages in your Go projects. | 570 | 
| [go-critic/go-critic](https://github.com/go-critic/go-critic) | The most opinionated Go source code linter for code audit. | 1247 | 
| [psampaz/go-mod-outdated](https://github.com/psampaz/go-mod-outdated) | Find outdated dependencies of your Go projects. go-mod-outdated provides a table view of the go list -u -m -json all command which lists all dependencies of a Go project and their available minor and patch updates. It also provides a way to filter indirect dependencies and dependencies without updates. | 564 | 
| [firstrow/go-outdated](https://github.com/firstrow/go-outdated) | Find outdated golang packages | 44 | 
| [yuroyoro/goast-viewer](https://github.com/yuroyoro/goast-viewer) | Golang AST visualizer | 597 | 
| [GoCover.io - GoCover.io offers the code coverage of any golang package as a service.](https://gocover.io/?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [goimports - Tool to fix (add, remove) your Go imports automatically.](https://godoc.org/golang.org/x/tools/cmd/goimports?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [arxdsilva/golang-ifood-sdk](https://github.com/arxdsilva/golang-ifood-sdk) | Golang Ifood API SDK  | 8 | 
| [segmentio/golines](https://github.com/segmentio/golines) | A golang formatter that fixes long lines | 352 | 
| [golang/lint](https://github.com/golang/lint) | [mirror] This is a linter for Go source code. (deprecated) | 3938 | 
| [Golint online - Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.](http://go-lint.appspot.com/?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [jfeliu007/goplantuml](https://github.com/jfeliu007/goplantuml) | PlantUML Class Diagram Generator for golang projects | 826 | 
| [goreturns - Adds zero-value return statements to match the func return types.](https://sourcegraph.com/github.com/sqs/goreturns?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [dominikh/go-tools](https://github.com/dominikh/go-tools) | Staticcheck - The advanced Go linter | 4542 | 
| [shurcooL/gostatus](https://github.com/shurcooL/gostatus) | A command line tool that shows the status of Go repositories. | 245 | 
| [surullabs/lint](https://github.com/surullabs/lint) | Run linters from Go code -  | 66 | 
| [z7zmey/php-parser](https://github.com/z7zmey/php-parser) | PHP parser written in Go | 856 | 
| [dominikh/go-tools](https://github.com/dominikh/go-tools) | Staticcheck - The advanced Go linter | 4542 | 
| [verygoodsoftwarenotvirus/blanket](https://github.com/verygoodsoftwarenotvirus/blanket) | MOVED TO GITLAB | 14 | 
| [augmentable-dev/tickgit](https://github.com/augmentable-dev/tickgit) | Manage your repository's TODOs, tickets and checklists as config in your codebase. | 277 | 
| [preslavmihaylov/todocheck](https://github.com/preslavmihaylov/todocheck) | A static code analyser for annotated TODO comments | 379 | 
| [mdempsky/unconvert](https://github.com/mdempsky/unconvert) | Remove unnecessary type conversions from Go source | 315 | 
| [dominikh/go-tools](https://github.com/dominikh/go-tools) | Staticcheck - The advanced Go linter | 4542 | 
| [mccoyst/validate](https://github.com/mccoyst/validate) | A Go package to automatically validate fields with tags | 60 | 


# Editor Plugins

| Repository | Description | Stars |
|:---|:---|:---:|
| [josa42/coc-go](https://github.com/josa42/coc-go) | Go language server extension using gopls for coc.nvim. | 435 | 
| [msyrus/vscode-go-doc](https://github.com/msyrus/vscode-go-doc) | An Microsoft Visual Code extension for Golang to print symbol definition to output | 5 | 
| [Go plugin for JetBrains IDEs - Go plugin for JetBrains IDEs.](https://plugins.jetbrains.com/plugin/9568-go?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [theia-ide/go-language-server](https://github.com/theia-ide/go-language-server) | A Go language server. | 31 | 
| [dominikh/go-mode.el](https://github.com/dominikh/go-mode.el) | Emacs mode for the Go programming language | 1227 | 
| [joefitzgerald/go-plus](https://github.com/joefitzgerald/go-plus) | An Enhanced Go Experience For The Atom Editor | 1513 | 
| [nsf/gocode](https://github.com/nsf/gocode) | An autocompletion daemon for the Go programming language | 4959 | 
| [incu6us/goimports-reviser](https://github.com/incu6us/goimports-reviser) | Right imports sorting & code formatting tool (goimports alternative) | 112 | 
| [goprofiling - This extension adds benchmark profiling support for the Go language to VS Code.](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego) |  | 0 | 
| [DisposaBoy/GoSublime](https://github.com/DisposaBoy/GoSublime) | A Golang plugin collection for SublimeText 3, providing code completion and other IDE-like features. | 3426 | 
| [hexdigest/gounit-vim](https://github.com/hexdigest/gounit-vim) | Vim plugin for https://github.com/hexdigest/gounit | 22 | 
| [theia-ide/theia-go-extension](https://github.com/theia-ide/theia-go-extension) | Theia Go Extension | 16 | 
| [rjohnsondev/vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) | Vim compiler plugin for Go (golang) | 87 | 
| [fatih/vim-go](https://github.com/fatih/vim-go) | Go development plugin for Vim | 14379 | 
| [golang/vscode-go](https://github.com/golang/vscode-go) | Go extension for Visual Studio Code | 2489 | 
| [eaburns/Watch](https://github.com/eaburns/Watch) | Watches for changes in a directory tree and reruns a command in an acme win or just on the terminal. | 190 | 


# Go Generate Tools

| Repository | Description | Stars |
|:---|:---|:---:|


# Go Tools

| Repository | Description | Stars |
|:---|:---|:---:|


# Software Packages

| Repository | Description | Stars |
|:---|:---|:---:|


# DevOps Tools

| Repository | Description | Stars |
|:---|:---|:---:|


# Other Software

| Repository | Description | Stars |
|:---|:---|:---:|



# ---

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.

