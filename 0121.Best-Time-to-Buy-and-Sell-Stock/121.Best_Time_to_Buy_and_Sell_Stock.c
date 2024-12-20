int maxProfit(int* prices, int pricesSize) {
    int minPrice = 1e9;
    int maxProfit = 0;

    for (int i = 0; i < pricesSize; i++) {
        minPrice = (prices[i] < minPrice)? prices[i]:minPrice;
        maxProfit = ((prices[i] - minPrice) > maxProfit)? (prices[i] - minPrice):maxProfit;
    }

    return maxProfit;
}