# Practical Common Lisp

## Syntax and Semantics

  In most programming languages, the language processor--whether an interepter or a compiler--operates as a black box:you shove a sequence of characters representing the text of a program into the black box,either executes the behaviors indicated or produces a compiled version of the program that will execute the behaviors when it's run.Common Lisp defines two black boxes, one that translates text into Lisp objects and another that implements the semantics of the language in terms of those objects. The first box is called reader, and the second is called the evaluator.

  Each black box defines one level of syntax. The reader defines how strings of characters can be translated into Lisp objects called `s-expressions`.The evaluator then defines a syntax of Lisp forms that can be built out of `s-expressios`.
>the syntax of s-expressions understood by the reader and the syntax of Lisp forms understood by the evaluator.

### S-expressions
The basic elements of s-expressions are lists and atoms.

- Lists: delimited by parentheses and can contain any number of whitespace-separated elements.
- Atoms: everything else.
  - Numbers:integers, ratios and floating;all rationals-integers and ratios-are represented internally in "simplified" form.`-2/8` and `-1/4` aren't distinct;On the other hand, 1.0 and 1.0d0 and 1 are different objects(different types).
  - Strings literals:enclosed in double quotes.
  - Symbols:Almost any character can appear in a name except whitespace,open and close parentheses, double and single quotes,backtick,comma,colon,semicolon,backslash,and vertical bar.the reader will read foo,Foo,and FOO as the same symbol:FOO.However,\f\o\o and |foo| will both be read as foo,which is a different object than the symbol FOO.

#### Naming conventions
- constant:+{name}+
- global:*{name}*

### S-expressions As Lisp Forms
All legal list forms start with a symbol,but three kinds of list forms are evaluated in three quite different ways.The evaluator must determine whether the symbol that starts the list is the name of a functioin,a macro,or a special operator.
- function call
- macro
- special operator

### Truth, Falsehood, and Equality
`T` means true,`NIL` means false and nil.
Four generic quality predicates:"Always use eql and use eq when possible"
- eq:tests for "object identity",(eq x x) can evaluate to either true or false depending on implements.
- eql:like `EQ` except that is also is guaranteed to condider two objects of the same class representing the same mumeric or character value to be equivalent.`(eql 1 1)` is guaranteed to be true and (eql 1 1.0) is guaranteed to be false.
- equal:loosens the discrimination of eql to consider lists,strings,bit vectors and pathnames,two data types equivalent if the have the same structure and contents.recursively,according to equal.
- equalp:similar to equal except it's even less discriminating.It considers two strings equivalent if the contain the same characters,ignoreing differences in case.Numbers are equivalent under equalp if they represent the same mathematical value.

## Functions
syntax:
```lisp
(defun name (parameters*)
  "Optional documentation string."
  body-form*)
```

### Optional Parameters
```lisp
(defun name (parameters* &optional (parameter [default value [default-not-supplied]])+)
  "documentation"?
  body-forms*)
```

height-customized is `T` if you pass a height otherwise use width value as height and height-customized is `NIL`.
```lisp
(defun cal-rectange(width &optional (height width height-customized))
  (list (* width height) height-customized))
(cal-rectange 10 9); (90 T)
(cal-rectange 10); (100 NIL)
```

### Rest Parameters
Lisp lets you include a catchall parameter after the symbol `&rest`.If a function includes a &rest parameter, any arguments remaining after values have been doled out to all the required and optional parameters are gathered up into a list that becomes the value of the &rest parameter.
```Lisp
(defun +me+(&rest values) (list values))
(+me+ 8 9 10) -> ((8 9 10))
```

### Keywords Parameters
Optional parameters are still positional--if the caller wants to pass an explicit value for the fourth optional parameter, it turns the first three optional parameters into required parameters for that caller.
To give a function keyword parameters,after any required,&optional,and &rest parameters you include the symbol `&key`.

```Lisp
(defun foo(a &optional(b "b-default") &rest values &key c)
  (list a b values c))

(foo 1 "b-value" 3 4 5 :c "c-value") -> error
(foo 1 "b-value" 3 :c "c-value") -> error
(foo 1 "b-value :c "c-value") -> (1 "b-value" (:C "c-value") "c-value")
```

If you want to completely decouple the public API of the function from the internal details, you can replace the parameter name with another list containering the keyword and optionally default value like `&optional`.
```Lisp
(defun foo(&key ((:apple a) 0) ((:box b) 0 not-default)) 
  (list a b not-default))

(foo :apple "a-value") -> ("a-value" 0 NIL)
(foo :box "b-value") -> (0 "b-value" T)
```

### Mixing Different Parameter Types
It's possible, but rare, to use all four flavors of parameters in a single function.Whenever more than one flavor of parameter is used, they must be declared in the order:
- names of the required parameters.
- the optional parameters
- the rest parameters|
- the keyword parameters

1.Combining &optional and &key parameters yields suprising enough results that you should probably avoid it altogether.

```Lisp
(defun foo(x &optional y &key z) (list x y z))
(foo 1 2 :z 3) -> (1 2 3)
(foo 1) -> (1 NIL NIL)
(foo 1 :z 3) -> Error(odd number of Key parameter arguments),:z is taken as optional parameter leaving only the argument 3 to be processed.
```

2.You can safely combine &rest and &key parameters,but the behavior may be a bit suprising.
All the remaining values, which include the keywords themselves,are gathered into a list that's bound to the &rest parameter, and the appropriate values are alse bound to the &key parameters.
```Lisp
(defun foo(&rest rest &key a b c)(list rest a b c))
(foo :a 1 :b 2 :c 3) -> ((:A 1 :B 2 :C 3) 1 2 3)
```

### Function Return Values
Returning the value of the last expression evaluated is the default behavior.However,you can use the `RETURN-FROM {block-name} {return-value}` special operator to immediately return any value from the function.

```Lisp
(defun foo(n)
  (dotimes (i 10)
    (dotimes (j 10)
      (when (> (* i j) n)
        (return-from foo (list i j))))))
```

## Numbers
One of the resones Lisp is a nice language for math is its numbers behave more like true mathematical numbers than the approximations of numbers that are easy to implement in finite computer hardware.Integers in Common Lisp can be almost arbitrarily large rather than being limitted by the size of a machine word.And dividing two integers results in an exact ratio,not a truncated value.

### Literals
The #B,#X,#O,#R syntaxes work only with rationals.

```Lisp
123 -> 123
+123 -> 123
123. -> 123
2/3 -> 2/3
#b10101 -> 21 ;; base2
#xa -> 10     ;; base16
#xA -> 10     ;; base16
#o777 -> 511  ;; base8
#5r4  -> 4    ;; base5
#b1010/1011 -> 10/11
```

## Characters
The Read syntax for characters objects is simple:#\ followed by the desired character.Whitespace,however, should be written as #\Space insteed of #\ ,Other names are Tab,Page,Rubout,Linefeed(Newline),Return,Newline,and Blackspace.

The case-sensitive analog to the numberic = is the function HCAR=.The case-insensitive version is CHAR-EQUAL.

### Character Comparison Functions.
| Numeric Analog | Case-Sensitive | Case-Insensitive |
| :-----         | :-----         | :-----           |
| =              | CHAR=          | CHAR-EQUAL       |
| /=             | CHAR/=         | CHAR-NOT-EQUAL   |
| <              | CHAR<          | CHAR-LESSP       |
| >              | CHAR>          | CHAR-GREATERP    |
| <=             | CHAR<=         | CHAR-NOT-GREATERP|
| >=             | CHAR>=         | CHAR-NOT-LESSP   |

## String
You can compare strings using a set of functions that follow the same naming convention as the character comparison functions except with STRING as the prefix rather than CHAR.However,unlike the character and number,the string comparators can compare only two strings and the arguments--:start1,:end1,:start2,:end2--specify the starting(inclusive) and ending(exclusive) indices of substrings.

```Lisp
(string= "foobarbaz" "quuxbarfoo" :start1 3 :end1 6 :start2 4 :end2 7)
(string= "foobarbaz" "quuxbarfoo" :start1 3 :end1 6)
```

## Collections

### Vectors
Vectors are basic integer-indexed collections in Common Lisp, and the came in two flavors(fixed-size vector and resizable vector).

#### create vectors
```Lisp
(vector) ->#()
(vector 1 2) ->#(1 2)

; MAKE-ARRAY is more general than VECTOR
; (MAKE-ARRAY n |'(n [m]) [:fill-pointer fill] [:initial-element initial] [:adjustable t] :element-type '{typename})
(make-array 2 :initial-element nil) -> #(NIL NIL)
(make-array '(2 3)) -> #2A((0 0 0) (0 0 0))
```

## Sequences functions
Vectors and lists are the two concrete subtypes of the abstract type `sequence`.

- lenght: return the length of a sequence. For vectors with a fill pointer, this will be the value of the fill pointer.
- elt:short for `element` and setable

```Lisp
(lenght *x*)
(elt *x* 0)
(setf (elt *x* 1) 99)
```

Basic Sequence Functions
| Name       | Required Arguments         | Returns |
| count      | Item and sequence          |         |
| find       | Item and sequence          |         |
| position   | Item and sequence          |         |
| remove     | Item and sequence          |         |
| substitute | New item,item,and sequence |         |

Standard Sequence Function Keyword Arguments
| Argument  | Meaning                                                    | Default |
| :test     |                                                            | eql     |
| :key      | function to extract key value from actual sequence element | nil     |
| :start    |                                                            | 0       |
| :end      |                                                            | nil     |
| :from-end | if traverse in reverse order from end to start.            | nil     |
| :count    | remove and substitute only                                 | nil     |

### Higheer-Order Function Variants
One set of variants are named the same as the basic function with an -IF or -IF-NOT sufix.
```Lisp
(count-if #'evenp #(1 2 3 4 5)) -> 2

(count-if-not #'evenp #(1 2 3 4 5)) -> 3

(remove-if-not #'(lambda (x) (char= (elt x 0) #\f)) 
    #("foo", "bar")) -> #("foo")
```

### Whole Sequence Manipulations
| Name        | Meaning | Arguments                                             |
| copy-seq    |         |                                                       |
| reverse     |         |                                                       |
| concatenate |         | Result type name,sequence a,and sequence b            |
| sort        |         | Sequence and comparator                               |
| merge       |         | Result type name,sequence a,sequence b,and comparator |
| map         |         | Result type name,function,and n sequences...          |
| map-into    |         | Result holder,function,and nsequences...              |

Example:
```Lisp
(sort (vector "foo", "bar") #'string<) -> #("bar", "foo")
(merge 'list #(1 3 5) #(2 4 6) #'<) -> (1 2 3 4 5 6)
(setf seq-a (vector 1 2 3))
(setf seq-b (vector 99 3))
(map 'vector #'+ (vector 1 2 3) #(2 3)) -> #(3 5)
(map-into seq-a #'+ seq-a seq-b) -> #(100 5 3)
```

### Subsequence Manipulations
```Lisp
; like list[begin:end] in the Python language,
; subsequence would never effect the raw sequence.
(subseq seq begin [end])
```

### Sequence Predicates
Four handy functions are `every`,`some`,`notany`,and`notevery`,which iterate over sequences testing a boolean predicate.({fn} {visitor-function} {seq}) -> boolean

## Hash Tables
```Lisp
(setf x (make-hash-table :test eql))
(setf (gethash 'name x) "akun")

; iter the table
(maphash #'(lambda (k v) (format t "~A=~A~%." k v) ))

; iter the table
(loop for k being the hash-keys in x using(hash-value v)
    do (format t "~A=~A~%" k v))
```

## List Processing
Historically,lists were Lisp's original composite data type.These days,a Common Lisp programmer is as likely to use a vector,a hash table,or a user-defined class or structure as to use a list.

### Cons Cells
CONS takes two arguments and returns a new cons cell containing the two values.Unless the second value is NIL or another cons cell,a cons isprinted as the two values in a parentheses separated by a dot,a so-called dotted pair.
```Lisp
(cons 1 nil) -> (1)
(cons 1 2) -> (1 . 2)
(car (cons 1 2)) -> 1
(cdr (cons 1 2)) -> 2
(cdr (cons 1 (cons 2 3))) -> (2 . 3)
```

- append: (append c1 c2), make a new copy of c1 and set it's latest cdr to c2

| Function          | Recycling Version | Description |
| :--               | :--               | :--         |
| append            | nconc             |             |
| remove            | delete            |             |
| remove-if         | delete-if         |             |
| remove-if-not     | delete-if-not     |             |
| remove-duplicates | delete=duplicates |             |
| substitute        | nsubstitute       |             |

## Object Oriented Programing

A non-object-oriented example:
```Lisp
;; define rectangle and circle struct
(defstruct rectangle
  height width)
(defstruct circle
  radius)

;; define area function
(defun area(x) 
  (cond((rectangle-p x) (* (rectangle-height x) (rectangle-width x))) 
       ((circle-p x) (expt (circle-radius x) 2)))) 

;; calculate area
(let ((r (make-rectangle :height 3 :width 9)))
  (area r))
```

A object-oriented example:
```Lisp
;; remove classes defined before
(setf (find-class 'rectangle) nil)
(setf (find-class 'circle) nil)
(setf (find-class 'area) nil)

;; define metaobjects
(defclass rectangle()
  (height width))
(defmethod area2((r rectangle))
  (* (slot-value r 'height) (slot-value r 'width)))

(defclass circle()
  (radius))
(defmethod area2((r circle))
  (expt (slot-value r 'radius) 2))

;; call area2 to calculate area of a circle
(let ((x (make-instance 'circle))) 
  (setf (slot-value x 'radius) 99)
  (area2 x))
```

How to define define a class:
```Lisp
(defclass name (superclasses)
    (slots)
    options)
```

A class define example:
```Lisp
(defclass Person()
  ((name 
        :writer set-person-name
        :reader get-person-name
        :initform "default Name"
        :initarg :name)
   (age 
        :accessor person-age
        :initform 0
        :initarg :age)
   (counter 
        :accessor person-counter
        :initform 0 
        :allocation :class
        :documentation "Describes the counter slot"))

   (:documentation "Describes a person"))

(setf v (make-instance 'Person))

;; slot setter
(setf (person-age v) "eage")
(set-person-name "akun" v)

;; slot getter
(get-person-name v)  ;;->"akun"
(person-age v)

;; type check
(type-of v) ;; ->PERSON
(class-of v) ;; ->::PERSON
(subtypep  (type-of v) 'Person) ;;->T

;; Implement the function `(is-instance v 'Person)`
(defun is-instance (v klass)
  (eq (find-class klass) (class-of v)))
(is-instance v 'Person) ;; ->T
```
