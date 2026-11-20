(def col1 [1 2 3 4 5 6 7 8])
(def col2 '(10 15 20 25 30 35))

(println "=== Collection 1 (Vector) ===")
(println "Original:" col1)
(println "map (square):          " (map #(* % %) col1))
(println "filter (evens):        " (filter even? col1))
(println "reduce (sum):          " (reduce + col1))
(println "reduce (max):          " (reduce max col1))
(println "take 4:                " (take 4 col1))
(println "drop 4:                " (drop 4 col1))
(println "Threading (->>) result:"
         (->> col1
              (filter even?)
              (map #(* % %))
              (take 3)
              (reduce +)))

(println "\n=== Collection 2 (List) ===")
(println "Original:" col2)
(println "map (double):          " (map #(* 2 %) col2))
(println "filter (> 20):         " (filter #(> % 20) col2))
(println "reduce (sum):          " (reduce + col2))
(println "reduce (max):          " (reduce max col2))
(println "take 3:                " (take 3 col2))
(println "drop 2:                " (drop 2 col2))
(println "Threading (->>) result:"
         (->> col2
              (drop 1)
              (filter even?)
              (map #(+ % 5))
              (reduce +)))
