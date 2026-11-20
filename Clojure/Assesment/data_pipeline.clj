
(ns data-pipeline
  (:require [ring.adapter.jetty :as jetty]
            [ring.util.response :as response])
  (:import [java.sql DriverManager SQLException]))

(def students
  [{:name "Adit" :marks 85}
   {:name "Rahul" :marks 32}
   {:name "Priya" :marks 76}
   {:name "Aman" :marks 45}
   {:name "Riya" :marks 28}])

(defn process-data []
  (let [passed (->> students
                    (map (fn [student]
                           (assoc student :passed
                                  (>= (:marks student) 40))))
                    (filter :passed))
        total (reduce + (map :marks passed))
        average (/ total (count passed))]
    {:students passed
     :average average}))

(def db-url "jdbc:sqlite:students.db")

(defn initialize-db []
  (try
    (with-open [connection (DriverManager/getConnection db-url)
                statement (.createStatement connection)]
      (.execute statement
                "CREATE TABLE IF NOT EXISTS students (name TEXT, marks INTEGER)"))
    (catch SQLException e
      (println "Database Error:" (.getMessage e)))))

(defn save-data [data]
  (try
    (with-open [connection (DriverManager/getConnection db-url)
                statement (.prepareStatement connection
                                             "INSERT INTO students (name, marks) VALUES (?, ?)")]
      (doseq [student data]
        (.setString statement 1 (:name student))
        (.setInt statement 2 (:marks student))
        (.executeUpdate statement)))
    (catch SQLException e
      (println "Database Error:" (.getMessage e)))))

(defn read-data []
  (try
    (with-open [connection (DriverManager/getConnection db-url)
                statement (.createStatement connection)
                result (.executeQuery statement "SELECT * FROM students")]
      (loop [rows []]
        (if (.next result)
          (recur (conj rows
                       {:name (.getString result "name")
                        :marks (.getInt result "marks")}))
          rows)))
    (catch SQLException e
      (println "Database Error:" (.getMessage e))
      [])))

(defn app [request]
  (if (= (:uri request) "/results")
    (response/content-type
     (response/response (pr-str (read-data)))
     "text/plain")
    (response/response "Welcome to Clojure Data Pipeline")))

(defn -main []
  (initialize-db)
  (save-data (:students (process-data)))
  (println "Server running at http://localhost:3000")
  (jetty/run-jetty app {:port 3000 :join? false}))
