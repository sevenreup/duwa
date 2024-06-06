mawu[] board = ["","","","","","","","",""];
mawu currentPlayer = "X";
nambala isRunning = zoona;

ndondomeko printBoard() {
    lembanzr("---------");
    lembanzr(board[0] + " | " + board[1] + " | " + board[2]);
    lembanzr("---------");
    lembanzr(board[3] + " | " + board[4] + " | " + board[5]);
    lembanzr("---------");
    lembanzr(board[6] + " | " + board[7] + " | " + board[8]);
    lembanzr("---------");
}

ndondomeko move(player, position) {
    board[position] = player;
    ngati (player == "X") {
        currentPlayer = "O";
    } kapena {
        currentPlayer = "X";
    }
    lembanzr(currentPlayer);
    lembanzr(player);
}

ndondomeko playGame() {
    lembanzr(isRunning);
    pamene(isRunning) {
        printBoard();
        nambala playerMove = kuNambala(console.landira());
        move(currentPlayer, playerMove - 1);
        // console.fufuta();
        lembanzr(isRunning);
    }
}

playGame();