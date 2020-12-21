cross(a,b).
cross(b,a).
move([X,X,Goat,Cabbage],wolf,[Y,Y,Goat,Cabbage]) :- cross(X,Y).
move([X,Wolf,X,Cabbage],goat,[Y,Wolf,Y,Cabbage]) :- cross(X,Y).
move([X,Wolf,Goat,X],cabbage,[Y,Wolf,Goat,Y]) :- cross(X,Y).
move([X,Wolf,Goat,Cabbage],nothing,[Y,Wolf,Goat,Cabbage]) :- cross(X,Y).

safe([X, Wolf, Ziege, Cabbage]) :- X = Ziege; Wolf = X, X = Cabbage.

solution([b,b,b,b], []).
solution(State, [Move|OtherMoves]) :- move(State,Move,NextState),
                                      safe(NextState),
                                      solution(NextState,OtherMoves).
run :- length(X,7), solution([a,a,a,a], X), write(X).