

(define (leaf-count tree)
  (if (pair? tree)
      (+ (leaf-count (car tree))
         (leaf-count (cdr tree)))
      1))


(define tree '((a . b) (c . d) . e))

(print (leaf-count tree))
