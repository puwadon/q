# q

 - quantum computation simulator


# linear algebra

```golang
v0 := vector.New(1, 1)
v1 := vector.New(1, -1)

v0.InnerProduct(v1) // -> complex(0, 0)
v0.IsOrthogonal(v1) // -> true
```

```golang
v0 := vector.New(1, 0)
v1 := v0.TensorProduct(v0) // -> Vector{1, 0, 0, 0}
```

```golang
x := gate.X() //-> X Gate of Pauli

v0 := vector.New(1, 0)
v0.Apply(x) // -> Vector{0, 1}

v1 := vector.New(1, 0, 0, 0)
x2 := x.TensorProduct(x)
v1.Apply(x2) // -> Vector{0, 0, 0, 1}
```

# quantum computation

## qubit

```golang
q := qubit.Zero()  // -> |0>
q.Apply(gate.H()) // -> 1/Sqrt(2) * (|0> + |1>)
q.Probability()   // -> (0.5, 0.5)
q.Measure()
q.Probability()   // -> (1, 0) or (0, 1)
```

## quantum circuit

```golang
# bell state
q := qubit.New(1, 0, 0, 0)
g0 := gate.H().TensorProduct(gate.I())
g1 := gate.CNOT()

bell := q.Apply(g0).Apply(g1)// -> 1/Sqrt(2) * (|00> + |11>)
// bell := q.Apply(g0.Apply(g1))

bell.Measure()
bell.Probability() // -> (1, 0, 0, 0) or (0, 0, 0, 1)
```

## quantum teleportation

```golang
g0 := gate.H().TensorProduct(gate.I())
g1 := gate.CNOT()
bell := qubit.Zero(2).Apply(g0).Apply(g1)

phi := qubit.New(1, 2) // arbitrary state
phi.Probability() // -> (0.2, 0.8)

phi.TensorProduct(bell)
g2 := gate.CNOT().TensorProduct(gate.I())
g3 := gate.H().TensorProduct(gate.I(2))
phi.Apply(g2).Apply(g3)

// Alice measure qubit 0, 1
mz := phi.Measure(0)
mx := phi.Measure(1)

// Alice send mx, mx to Bob
// Bob Apply Z and X
if mz.IsOne() {
  z := gate.I(2).TensorProduct(gate.Z())
  phi.Apply(z)
}

if mx.IsOne() {
  x := gate.I(2).TensorProduct(gate.X())
  phi.Apply(x)
}

// Bob got phi state
phi.Probability()
// One of the following:
// (0.2, 0.8, 0, 0, 0, 0, 0, 0)
// (0, 0, 0.2, 0.8, 0, 0, 0, 0)
// (0, 0, 0, 0, 0.2, 0.8, 0, 0)
// (0, 0, 0, 0, 0, 0, 0.2, 0.8)
```

## Grover's search algorithm

```golang
x := matrix.TensorProduct(gate.X(), gate.I(3))
oracle := x.Apply(gate.CNOT(4)).Apply(x)

h4 := matrix.TensorProduct(gate.H(3), gate.H())
x3 := matrix.TensorProduct(gate.X(3), gate.I())
cz := matrix.TensorProduct(gate.CZ(3), gate.I())
h3 := matrix.TensorProduct(gate.H(3), gate.I())
amp := h4.Apply(x3).Apply(cz).Apply(x3).Apply(h3)

q := qubit.TensorProduct(qubit.Zero(3), qubit.One())
q.Apply(gate.H(4)).Apply(oracle).Apply(amp)

q.Probability()
// [0 0.03125 0 0.03125 0 0.03125 0 0.78125 0 0.03125 0 0.03125 0 0.03125 0 0.03125]
```

# Reference

 1. Michael A. Nielsen, Issac L. Chuang, Quantum Computation and Quantum Information
 2. C. Figgatt, D. Maslov, K. A. Landsman, N. M. Linke, S. Debnath, and C. Monroe, Complete 3-Qubit Grover Search on a Programmable Quantum Computer
