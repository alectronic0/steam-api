# Steam Api

This is a simple Library to access Steam Public API's

With a proof of concept comparing two people steam library to show what they have in common.

## How to run
The initial setup require creating an `.env` to store your Steam Dev API Key & The 2 user ID you want to compare

The API's require the usage of public data (for you and other users)
If you want to see this information you will need to adjust your setting here: https://steamcommunity.com/id/me/edit/settings


### Create a Dev Key:
To Create a Dev API Key go here (Do not share with anyone): https://steamcommunity.com/dev/apikey

T&C: https://steamcommunity.com/dev/apiterms

You are Limited to 100,000 Per Day (So don't go crazy)

### Find Steam ID
The API uses the steamID64 to be exact there are a various number of way to find this however here are a few ways:
- Your Steam ID can be found on you account page: https://store.steampowered.com/account
- You can find other by searching in https://www.steamidfinder.com
- Coming soon: there is an existing API to pass in `/ISteamUser/ResolveVanityURL/v0001/` however it on the backlog to do.

### Create .env File
The below command will setup the `.env` with the required field please add the value to gather earlier
```bash
touch .env
echo "STEAM_API_KEY=" >> .env
echo "TEST_USER_1=" >> .env
echo "TEST_USER_2=" >> .env
```

### run stream comparator

```bash
 go run ./cmd/compare_steam_libs/main.go
```
this will create a `comparison_result.json` for you to view.

## Coming soon
- expose more steam client API fields:
  - User wishlist
  - Video game metadata
      - Game Stats
      - Store Details
- Build HTTP Web Server
  - with a redis cache of individual client calls
  - with docker image 
  - instruction to deploy
- Build Simple React Application
- Write Unit Test
