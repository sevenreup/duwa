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
}

ndondomeko playGame() {
    lembanzr(isRunning);
    pamene(isRunning) {
        printBoard();
        nambala move = console.landira();
        lembanzr(move);
        move(currentPlayer, move);
        // console.fufuta();
    }
}

playGame();