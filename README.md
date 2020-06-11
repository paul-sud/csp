# Communicating sequential processes in Go

Solutions to problems posited by C.A.R. Hoare in _Communicating_ _Sequntial_ _Processes_. See the PDF at https://www.cs.cmu.edu/~crary/819-f09/Hoare78.pdf . These are not necessarily the same solutions proposed in the paper, particulary the use of guards in a lot of the original solutions implies using Go `select`s over channels, which requires additional channels than the solutions I have.
