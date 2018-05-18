

(define (slow_abs x)
  (if (>= x 0)
      x
      (- x)))

(print (slow_abs 3))
(print (slow_abs -3))
