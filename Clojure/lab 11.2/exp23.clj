(defn safe-divide [a b]
  (try
    (/ a b)
    (catch ArithmeticException e
      (println "Caught ArithmeticException:" (ex-message e))
      :division-error)
    (finally
      (println "Finally block executed for safe-divide"))))

(defn safe-string-length [s]
  (try
    (.length s)
    (catch NullPointerException e
      (println "Caught NullPointerException:" (ex-message e))
      0)
    (finally
      (println "Finally block executed for safe-string-length"))))

(defn withdraw [balance amount]
  (if (> amount balance)
    (throw (ex-info "Insufficient funds"
                    {:error-code :INSUFFICIENT_FUNDS
                     :balance balance
                     :requested amount
                     :shortfall (- amount balance)}))
    (- balance amount)))

(defn safe-withdraw [balance amount]
  (try
    (let [new-balance (withdraw balance amount)]
      (println "Withdrawal successful. New balance:" new-balance)
      new-balance)
    (catch clojure.lang.ExceptionInfo e
      (println "Caught custom ExceptionInfo:")
      (println "  ex-message:" (ex-message e))
      (println "  ex-data:   " (ex-data e))
      (println "  Recovering gracefully -> keeping original balance:" (:balance (ex-data e)))
      (:balance (ex-data e)))
    (finally
      (println "Finally block executed for safe-withdraw"))))

(println "=== (1) ArithmeticException ===")
(println "Result:" (safe-divide 10 0))

(println "\n=== (2) NullPointerException ===")
(println "Result:" (safe-string-length nil))

(println "\n=== (3) Custom ex-info & Graceful Recovery ===")
(println "Result (valid):  " (safe-withdraw 500 200))
(println "Result (invalid):" (safe-withdraw 500 800))
