## Question - Merkle Tree

1) In the illustration, let’s imagine I know the whole Merkle tree. Someone gives me L2 data block but I don’t trust him. How can I check if L2 data is valid?
-> If I know the whole tree I would need to get the sibling of Hash 0-1(L2), which is Hash 0-0(L1), then concatenate them together in order to produce a hash. Then compare this hash with Hash 0. If same then L2 data received is valid.

2) I know only the L3 block and the Merkle root. What minimum information do I need to check that the L3 block and the Merkle root belong to the same Merkle tree ?
-> I need the sibling of Hash 1-0 (L3), which is Hash 1-1 (L4) and their parent's sibling Hash 0.

3) What are some Merkle tree use cases?
-> In systems where data integrity needs to be done: Distributed systems most of all but not only, ie: database replication.

## Question - Modelisation Task

1. What would be an efficient way to modelize this? (models and database)

In a realtime DB (NoSQL) like RethinkDb:

- 1 Table with all cards of the game and their characteristics (owner, playing, oldOwners, performance ...).
```json
{ card.json
  "name": "Zinédine Zidane",
  "type": {
    "Unique": {
      "cards": [{
          "id": "0x-Zizou",
          "isPlaying": true, // is already playing in tournament
          "attack": 100,
          "vista": 100,
          "listOwner": ["PlayerName", "..."]
      }]
    },
    "Common": {
      "cards": [{
          "id": "1x-Zizou",
          "attack": 90,
          "vista": 90,
          "listOwner": ["PlayerName", "..."]
      }]
    }
  }
}
```

- 1 Table with all players account with their infos (name, list of cards, palmares ...).
```json
{ player.json
  "name": "Arsene Wenger",
  "cards": ["0x-Mbappe", "0x-Zizou"],
  "palmares": {
    "World Cup": 1
  }
}
```

- 1 Table tournaments that gets "renewed" every Monday and is readonly only from Tuesday to Thursday.
    -> Each league is a JSON object looking like:
```json
{ league.json
  "name": "Ligue 1",
  "cards": {
    "0x-Zizou": ["Arsene Wenger", ["..."]]
   },
  "players": {
    "Arsene Wenger": {
      "cards": ["0x-Zizou", "1x-Zizou", "2x-Zizou", "3x-Zizou", "4x-Zizou"],
      "creationTime": 123432343,
      "score": 100
    }
  }
}
```

The field "cards" allows us to modify teams score efficiently. If a player stats needs to be update and therefore changing the score of its team. Instead of having to go through all teams to see which team has it, we can keep the track of all the cards used and by which player. Allowing us to update the data more efficiently.

2. What is your strategy to update rankings in real time?
Our backend is connected to our frontend via Websocket. We have a function listening to changes in our db making it "realtime". When "things" happens and the db get updated (let say Zidane is playing and got a score of 100 during the weekend) then our backend will properly send this info to all our users through websocket. Update will be "live" on the frontend.

3. What are the trade-offs?
Possibility of big json objects. Possibility of sending update to users that don't need it. Needs of good replication in our databases (rethinkdb handles it but in the case where we would build our own system we could use a merkle tree). Not thinking of anything else at the moment.