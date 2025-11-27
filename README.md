# goContexts

## understanding go contexts in a better way.
### Topics discussed 
1) ctx, cancel := context.WithCancel(context.Background())
2) ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
3) time.AfterFunc(2*time.Second, cancel)
4) ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(0*time.Second))
5) ctx := context.WithValue(context.Background(), nestedroutines.AuthParamKey, authParamKeyVal)
6) checking context within http request
