# Curated RSS Feed

## Install using Go

You'll need [Go](https://golang.org) and either [Reeder](https://reederapp.com) or [Cappuccino](http://cappuccinoapp.com).

```sh
go get -u github.com/cixtor/rssfeed
```

## Install using Docker

You'll need [Docker Desktop](https://www.docker.com/products/docker-desktop) and either [Reeder](https://reederapp.com) or [Cappuccino](http://cappuccinoapp.com).

```sh
docker-compose up -d
```

## Usage

* Request an API token to communicate with [Mercury web API service](https://mercury.postlight.com/web-parser/)
* Add the token visiting http://localhost:9628/register?token=API_TOKEN
* Use your favorite RSS client to subscribe to http://localhost:9628/news.rss

| Reeder | Cappuccino |
|--------|------------|
| <img src="images/screenshot1.png" width="418"> | <img src="images/screenshot2.png" width="418"> |

## Story Time

I'm a fan of RSS, too bad many people are killing it, directly or indirectly.

I few months ago I discovered a new app for macOS called [Cappuccino](http://cappuccinoapp.com). Very nice —native— interface, minimalistic. Unfortunately, being a new app, sooner than later I started to find bugs here and there, not with the interface but with the content that many websites were returning in their RSS feeds, some of which would only render a title and a link to view the actual article _(looking at you Hacker News)_.

Fast-foward a few days, [Reeder](https://reederapp.com) —another famous macOS app— got released for free in the Mac App Store. I immediately switched and imported my OPML file. Working pretty good, I have to say, but I was still facing the same problems that I had with Cappuccino: many RSS feeds were returning just an excerpt of the actual article. Fortunately Reeder offers a function called "Mercurity Reader" where they send the link to a web API service —owned by [Postlight](https://mercury.postlight.com)— and returns a clean version of the content, similar to what Safari, Firefox and others do with the "Reader View".

Unfortunately, this forces me to click 1-2 more times to start reading.

Wait! Am I really complaining about 1-2 extra clicks?

Well, yes, what's the point of RSS if you cannot read the content from there?

I understand that in today's world of "News by Subscription" adding the full content of the article to the RSS feed allows people to pirate the content, read "piracy" as _"go around the paywall"_. But if your website wasn't so bloated with JavaScript and ads I would surely visit it and pay for the subscription.

I immediately started working on a web service to act as a proxy for all the RSS feeds that I am subscribed to. The API would take all the links from these feeds and send them to Mercuriy, cache them locally, and then returning a bigger RSS feed _(with full content)_ to either Cappuccino or Reeder.

## Fallback Mechanism

Some news websites block Mercury, so I had to implement a fallback mechanism to parse the website if Mercury is not able to access its content. I did this with simple string replacements. The code is quite flacky, and will surely break, but if it's just me sending requests then maybe they will notice.
