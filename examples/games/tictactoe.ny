ndondomeko printBoard() {
    console.lemba(" 1 | 2 | 3 ");
    za (nambala i = 0; i < 3; i++) {
        console.lemba("---------");
        console.lemba(" 4 | 5 | 6 ");
        console.lemba("---------");
        console.lemba(" 7 | 8 | 9 ");
    }
}

ndondomeko playGame() {
    mawu currentPlayer = "X";
    nambala isRunning = zoona;
    console.lemba(isRunning);
    pamene(isRunning) {
        printBoard();
        isRunning = bodza;
    }
}

playGame();