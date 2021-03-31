# Go-GA

Golang Genetic Algorithm with multithread support

## Algorithm

*Pseudo-code*
```
   Initialize Population
   Calculate Fitness of Population
   
   while (termination criteria isnt met) do
      Parent Selection
      Crossover with probability Pc
      Mutation with probability Pm
      Calculate Fitness
      Survivor Selection
      Select Best
   return Best
```
