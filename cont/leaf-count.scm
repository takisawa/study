
(define tree '((a . b) (c . d) . e))

(define (leaf-count tree)
  (if (pair? tree)
      (+ (leaf-count (car tree))
         (leaf-count (cdr tree)))
      1))


(define (leaf-count/cps tree cont)
  (if (pair? tree)
      (leaf-count/cps (car tree)
          (lambda (n)
            (leaf-count/cps (cdr tree)
              (lambda (m) (cont (+ n m))))))
      (cont 1)))

(print (leaf-count tree))
(print (leaf-count/cps tree identity))
