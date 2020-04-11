# wrigleyBall
A simulation of a dice simulation of a little league like game.

Wrigley Ball is Copyright Mike Forsberg 2020
Rules and Source Code Licensed under GNU GPL v3,

If you want changes:
- Please fork the repo, make the changes, and issue a pull request

## Prologue

As said by the tabaxi Tower to the Sky, a 8th level bard, *“Wrigley Ball is a game for commoners, being played by demi-gods”.*

## Hardware
- 2 3x3 tiles for difficult terrain
- 1 1x4 tile for difficult terrain
- 1 5x5 tile for difficult terrain (for teams of 4-5 players)
- 1 home team of 1-5 D&D characters
- 1 away team of 1-5 D&D characters
- 1 10x10 grid for teams of 1-3 players; 1 20x20 grid for teams of 4-5 players
- A score board to keep track of score and which team is batting for a giving turn.

## Game board description

The game board is addressed with percentile dice or 2d20s.
The tens die or first d20 is the north-south coordinate.
The ones die or second d20 is the east-west coordinate.
Numbering starts in the lower left square.

## Game board setup

Before the game is started the difficult terrain tiles are placed.
All tiles are placed by rolling a d100 or 2d20s
The lower left corner of the tile is placed on the specific square rolled.
For the 1x4 tile, a d6 is rolled and the tile is placed with the longer dimension aligned north-south if the value is even, east-west if the value is odd.
Difficult terrain may lie off the game board or overlap other difficult terrain due to certain rolls.

## Batting order.

The batting order for the duration of the game is assigned by the team as they wish.
Once created, the batting order is fixed for the duration of the game.
2 players, one from each team, form a batting/pitching pair.
Thus the first positions from each team are paired up and bat and pitch together. The second position players pair up and so forth.
If one of the players becomes inactive, the player does not bat and another team member must pinch pitch for that player.
Pinch pitching does not affect the batting order.
If both players of the batting/pitching pair are inactive, that pair is skipped.
The away team always batts first in the pair order.

## Before the first pitch

Before pitching is started, starting with the second in the batting order, the players arrange themselves around the field.
Players alternate turns starting with the away team following second in the batting order and place themselves on the field except for the squares pitching mound and batter’s box.
On a 10x10 grid the pitching mound is 77 and the batter’s box is 33.
On a 20x20 grid the pitching mound is 13,13 and the batter’s box is 5,5.
The players first in the batting order are placed on the pitching mound and batter’s box respective to their future rolls.

## Before each pitch

The current pitcher and batter are placed on the correct location prior to the pitch.  They may move from their locations once the balls are in play.  All other players do not move before pitching and batting.

## Playing Wriggley ball

After the pitcher and batter players are moved from their current location on the field to the respective batting box and pitching mound.  The pitching may begin.

### Rolling for a pitch
The pitcher rolls **1d12** while the batter rolls a **2d6**.  If the pitcher rolls higher than the batter, the batter is out.  If the batter ties or rolls higher than the pitcher, the difference between the batter's roll and pitcher's roll places that many balls **+1** into play.  For instance:

|Batter|Pitcher|Num of balls in play|
| --- | --- | --- |
|10|12|Batter is out|
|10|5|6 balls in play (5+1) |
|2|2|1 ball in play (0+1)|

## Placing balls in play
For each ball placed into play a d100 or 2d20s are rolled and the ball is placed on that square. 

If the ball lands on a square occupied by a pitching team player, the ball is immediately destroyed. (No action needed)

If the ball lands on a square occupied by a hitting team player, the automatically counts as a score. (No action needed)

After all the balls are placed into play or have been scored and destroyed, the players may move, act on the balls, or attack each other for one round of initiative.  

The ordering of action after the balls are hit is based on a fresh initiative roll by each of the active players.

## To act on a ball

If the player enters the square of a ball, the player may choose to use their action to either:
- If the player's team is currently hitting, remove the ball from play and count it as a score.
- If the player's team is currently pitching, remove the ball from play.  Denying the hitting team from the ability to score.

If the player attacks a ball at range, the AC of the ball is 13.
The HP of the ball is 1.
If the player’s team is currently pitching, successfully attacking the ball at range destroys the ball.
If the player’s team is currently hitting, successfully attacking the ball at range counts the ball as a score and destroys the ball.
Only one ball may be attacked at range per a player’s turn.

**Destroying/Denying and scoring on a ball is an action.**

## Play continues
Play continues in initiative order until all balls are removed from the field either by scoring or destroying/denying them.
After all balls are removed from play; all players stay in their current locations except for the current or next batting pair.  
After the away team hits the home team becomes the batter and the opposing batter now becomes the pitcher.
If the home team has finished batting, then the next pair in the batting order becomes the batter/pitcher pair.
After all remaining pairs have completed a round of batting, the next inning starts.
A total of three innings are allowed.
A game time of 1 minute elapses between each inning, but players do not move during this time. 

*Some innings are fast, others are long.  The choice of three innings is to limit the real world time to something manageable in one session*

## Attacking opposing players
 
All moves and actions that are permitted in DnD 5e are permitted during play,  When a player hits zero hit points, they become inactive and are removed from the field of play.  An inactive player can not bat, and must be pinched pitched.  

## What stops the game

The game is played for 3 innings, or 3 cycles of completing all batting orders.  The game is also halted if all players of one team are eliminated.  The score at the time of the last player from a team becoming eliminated is the final score.

## Winning the game

The only way to win the game is to score more points then the other team.  Even if the one team has had all their players eliminated, they can still win by having more points when their team is eliminated.
