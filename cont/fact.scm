
(define (fact n)
  (if (= n 1)
      1
      (* n (fact (- n 1)))))

(define (fact/cps n cont)
  (if (= n 1)
      (cont 1)
      (fact/cps (- n 1) (lambda (a) (cont (* n a))))))

(print (fact 5))
(print (fact/cps 5 identity))
