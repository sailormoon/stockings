# 🧦 stockings
Public domain equity tracker that fits in your tmux pane. Very rough around the edges in its current state and gets the bare minimum done. Use at your own risk.

# motivation
I want to glance at my portfolio performance in the corner of my terminal instead of having a browser tab open or checking my phone. It is also a good excuse to learn golang.

# usage
```
./stockings tickers
```

## example
The below symbols and quotes will continuously update.
```
./stockings mtch fb aapl mu msft
MTCH $36.27
  FB $174.02
AAPL $176.89
  MU $46.62
MSFT $94.07
```

# data
Data is provided for free from [IEX](https://www.iextrading.com/developer). Please read the following before contributing to this application or using it:
1. [API Terms of Use](https://iextrading.com/api-terms/)
2. [API Exhibit A](https://iextrading.com/api-exhibit-a/)
