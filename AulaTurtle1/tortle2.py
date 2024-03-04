import turtle as tr
import turtle as T

tr.pensize(5)
tr.speed(13)

#telhado
T.begin_fill()
T.forward(200)
T.left(120)
T.forward(200)
T.left(120)
T.forward(200)
T.fillcolor("brown")
T.end_fill()

#paredes/chao
T.begin_fill()
T.left(120)
T.forward(20)

T.right(90)
T.forward(160)

T.left(90)
T.forward(160)
T.left(90)
T.forward(160)
T.left(90)

T.fillcolor("yellow")
T.end_fill()

#porta
T.left(90)
T.forward(160)
T.right(90)
T.forward(140)

T.right(90)
T.right(180)
T.begin_fill()
T.left(90)
T.forward(80)
T.left(90)
T.forward(80)
T.left(90)
T.forward(40)
T.left(90)
T.forward(80)
T.fillcolor("brown")
T.end_fill()

#porta - detalhes
T.left(90)
T.forward(10)
T.left(90)
T.forward(80)
T.left(90)
T.forward(10)
T.left(90)
T.forward(80)
T.left(90)
T.forward(30)
T.left(90)
T.forward(80)
T.left(90)
T.forward(30)
T.left(90)
T.forward(80)
T.left(90)

#janelas
T.forward(100)
T.left(90)
T.forward(160)

T.left(90)
T.forward(160)
T.left(180)
T.forward(20)

tr.color('yellow')
T.right(90)
T.forward(20)

T.begin_fill()
tr.color('white')
T.forward(40)
T.left(90)
T.forward(120)
T.left(90)
T.forward(40)
T.left(90)
T.forward(120)
T.fillcolor("black")
T.end_fill()

T.left(90)
T.forward(40)
T.left(90)
T.forward(40)
T.left(90)
T.forward(40)
