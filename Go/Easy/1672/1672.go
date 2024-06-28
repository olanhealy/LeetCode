func maximumWealth(accounts [][]int) int {
    var maxWealth int

    for _, customer := range accounts {
        currentWealth := 0

        for _, bank := range customer {
            currentWealth  += bank
        }

        if (currentWealth > maxWealth ) {
            maxWealth = currentWealth
        }

    }

    

    return maxWealth

    
}
