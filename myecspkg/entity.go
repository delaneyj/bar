package myecspkg

import comp "bar/components"

const entityPageBits = 10
const entityPageSize = 1 << entityPageBits

type EntityID uint64

type ComponentMapping [3]uint64

type Entity struct {
	id         EntityID
	generation uint64
	components ComponentMapping
}

var freeList []EntityID
var entities [][]Entity

func (e Entity) Kill() {
    if !e.Alive() {
        return
    }
	freeList = append(freeList, e.id)
	entities[e.id >> entityPageBits][e.id % entityPageSize].generation++
	entities[e.id >> entityPageBits][e.id % entityPageSize].components = ComponentMapping{}
}

func (e Entity) Alive() bool {
    return e.generation == entities[e.id >> entityPageBits][e.id % entityPageSize].generation
}


func (e Entity) Position() *comp.Position {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 1 != 0 {
        return &storePosition[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) Velocity() *comp.Velocity {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 2 != 0 {
        return &storeVelocity[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) A() *comp.A {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 4 != 0 {
        return &storeA[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) B() *comp.B {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 8 != 0 {
        return &storeB[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) C() *comp.C {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 16 != 0 {
        return &storeC[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) D() *comp.D {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 32 != 0 {
        return &storeD[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) E() *comp.E {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 64 != 0 {
        return &storeE[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) F() *comp.F {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 128 != 0 {
        return &storeF[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) G() *comp.G {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 256 != 0 {
        return &storeG[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) H() *comp.H {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 512 != 0 {
        return &storeH[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) I() *comp.I {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 1024 != 0 {
        return &storeI[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) J() *comp.J {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 2048 != 0 {
        return &storeJ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) K() *comp.K {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 4096 != 0 {
        return &storeK[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) L() *comp.L {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 8192 != 0 {
        return &storeL[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) M() *comp.M {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 16384 != 0 {
        return &storeM[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) N() *comp.N {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 32768 != 0 {
        return &storeN[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) O() *comp.O {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 65536 != 0 {
        return &storeO[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) P() *comp.P {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 131072 != 0 {
        return &storeP[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) Q() *comp.Q {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 262144 != 0 {
        return &storeQ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) R() *comp.R {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 524288 != 0 {
        return &storeR[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) S() *comp.S {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 1048576 != 0 {
        return &storeS[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) T() *comp.T {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 2097152 != 0 {
        return &storeT[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) U() *comp.U {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 4194304 != 0 {
        return &storeU[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) V() *comp.V {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 8388608 != 0 {
        return &storeV[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) W() *comp.W {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 16777216 != 0 {
        return &storeW[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) X() *comp.X {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 33554432 != 0 {
        return &storeX[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) Y() *comp.Y {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 67108864 != 0 {
        return &storeY[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) Z() *comp.Z {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 134217728 != 0 {
        return &storeZ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AA() *comp.AA {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 268435456 != 0 {
        return &storeAA[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AB() *comp.AB {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 536870912 != 0 {
        return &storeAB[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AC() *comp.AC {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 1073741824 != 0 {
        return &storeAC[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AD() *comp.AD {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 2147483648 != 0 {
        return &storeAD[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AE() *comp.AE {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 4294967296 != 0 {
        return &storeAE[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AF() *comp.AF {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 8589934592 != 0 {
        return &storeAF[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AG() *comp.AG {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 17179869184 != 0 {
        return &storeAG[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AH() *comp.AH {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 34359738368 != 0 {
        return &storeAH[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AI() *comp.AI {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 68719476736 != 0 {
        return &storeAI[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AJ() *comp.AJ {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 137438953472 != 0 {
        return &storeAJ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AK() *comp.AK {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 274877906944 != 0 {
        return &storeAK[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AL() *comp.AL {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 549755813888 != 0 {
        return &storeAL[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AM() *comp.AM {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 1099511627776 != 0 {
        return &storeAM[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AN() *comp.AN {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 2199023255552 != 0 {
        return &storeAN[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AO() *comp.AO {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 4398046511104 != 0 {
        return &storeAO[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AP() *comp.AP {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 8796093022208 != 0 {
        return &storeAP[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AQ() *comp.AQ {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 17592186044416 != 0 {
        return &storeAQ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AR() *comp.AR {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 35184372088832 != 0 {
        return &storeAR[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AS() *comp.AS {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 70368744177664 != 0 {
        return &storeAS[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AT() *comp.AT {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 140737488355328 != 0 {
        return &storeAT[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AU() *comp.AU {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 281474976710656 != 0 {
        return &storeAU[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AV() *comp.AV {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 562949953421312 != 0 {
        return &storeAV[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AW() *comp.AW {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 1125899906842624 != 0 {
        return &storeAW[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AX() *comp.AX {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 2251799813685248 != 0 {
        return &storeAX[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AY() *comp.AY {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 4503599627370496 != 0 {
        return &storeAY[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) AZ() *comp.AZ {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 9007199254740992 != 0 {
        return &storeAZ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BA() *comp.BA {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 18014398509481984 != 0 {
        return &storeBA[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BB() *comp.BB {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 36028797018963968 != 0 {
        return &storeBB[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BC() *comp.BC {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 72057594037927936 != 0 {
        return &storeBC[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BD() *comp.BD {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 144115188075855872 != 0 {
        return &storeBD[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BE() *comp.BE {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 288230376151711744 != 0 {
        return &storeBE[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BF() *comp.BF {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 576460752303423488 != 0 {
        return &storeBF[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BG() *comp.BG {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 1152921504606846976 != 0 {
        return &storeBG[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BH() *comp.BH {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 2305843009213693952 != 0 {
        return &storeBH[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BI() *comp.BI {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & 4611686018427387904 != 0 {
        return &storeBI[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BJ() *comp.BJ {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[0] & -9223372036854775808 != 0 {
        return &storeBJ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BK() *comp.BK {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 1 != 0 {
        return &storeBK[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BL() *comp.BL {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 2 != 0 {
        return &storeBL[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BM() *comp.BM {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 4 != 0 {
        return &storeBM[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BN() *comp.BN {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 8 != 0 {
        return &storeBN[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BO() *comp.BO {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 16 != 0 {
        return &storeBO[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BP() *comp.BP {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 32 != 0 {
        return &storeBP[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BQ() *comp.BQ {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 64 != 0 {
        return &storeBQ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BR() *comp.BR {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 128 != 0 {
        return &storeBR[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BS() *comp.BS {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 256 != 0 {
        return &storeBS[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BT() *comp.BT {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 512 != 0 {
        return &storeBT[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BU() *comp.BU {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 1024 != 0 {
        return &storeBU[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BV() *comp.BV {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 2048 != 0 {
        return &storeBV[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BW() *comp.BW {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 4096 != 0 {
        return &storeBW[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BX() *comp.BX {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 8192 != 0 {
        return &storeBX[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BY() *comp.BY {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 16384 != 0 {
        return &storeBY[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) BZ() *comp.BZ {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 32768 != 0 {
        return &storeBZ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CA() *comp.CA {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 65536 != 0 {
        return &storeCA[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CB() *comp.CB {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 131072 != 0 {
        return &storeCB[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CC() *comp.CC {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 262144 != 0 {
        return &storeCC[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CD() *comp.CD {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 524288 != 0 {
        return &storeCD[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CE() *comp.CE {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 1048576 != 0 {
        return &storeCE[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CF() *comp.CF {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 2097152 != 0 {
        return &storeCF[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CG() *comp.CG {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 4194304 != 0 {
        return &storeCG[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CH() *comp.CH {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 8388608 != 0 {
        return &storeCH[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CI() *comp.CI {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 16777216 != 0 {
        return &storeCI[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CJ() *comp.CJ {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 33554432 != 0 {
        return &storeCJ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CK() *comp.CK {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 67108864 != 0 {
        return &storeCK[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CL() *comp.CL {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 134217728 != 0 {
        return &storeCL[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CM() *comp.CM {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 268435456 != 0 {
        return &storeCM[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CN() *comp.CN {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 536870912 != 0 {
        return &storeCN[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CO() *comp.CO {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 1073741824 != 0 {
        return &storeCO[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CP() *comp.CP {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 2147483648 != 0 {
        return &storeCP[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CQ() *comp.CQ {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 4294967296 != 0 {
        return &storeCQ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CR() *comp.CR {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 8589934592 != 0 {
        return &storeCR[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CS() *comp.CS {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 17179869184 != 0 {
        return &storeCS[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CT() *comp.CT {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 34359738368 != 0 {
        return &storeCT[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CU() *comp.CU {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 68719476736 != 0 {
        return &storeCU[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CV() *comp.CV {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 137438953472 != 0 {
        return &storeCV[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CW() *comp.CW {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 274877906944 != 0 {
        return &storeCW[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CX() *comp.CX {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 549755813888 != 0 {
        return &storeCX[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CY() *comp.CY {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 1099511627776 != 0 {
        return &storeCY[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) CZ() *comp.CZ {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 2199023255552 != 0 {
        return &storeCZ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DA() *comp.DA {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 4398046511104 != 0 {
        return &storeDA[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DB() *comp.DB {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 8796093022208 != 0 {
        return &storeDB[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DC() *comp.DC {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 17592186044416 != 0 {
        return &storeDC[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DD() *comp.DD {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 35184372088832 != 0 {
        return &storeDD[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DE() *comp.DE {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 70368744177664 != 0 {
        return &storeDE[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DF() *comp.DF {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 140737488355328 != 0 {
        return &storeDF[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DG() *comp.DG {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 281474976710656 != 0 {
        return &storeDG[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DH() *comp.DH {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 562949953421312 != 0 {
        return &storeDH[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DI() *comp.DI {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 1125899906842624 != 0 {
        return &storeDI[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DJ() *comp.DJ {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 2251799813685248 != 0 {
        return &storeDJ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DK() *comp.DK {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 4503599627370496 != 0 {
        return &storeDK[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DL() *comp.DL {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 9007199254740992 != 0 {
        return &storeDL[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DM() *comp.DM {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 18014398509481984 != 0 {
        return &storeDM[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DN() *comp.DN {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 36028797018963968 != 0 {
        return &storeDN[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DO() *comp.DO {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 72057594037927936 != 0 {
        return &storeDO[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DP() *comp.DP {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 144115188075855872 != 0 {
        return &storeDP[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DQ() *comp.DQ {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 288230376151711744 != 0 {
        return &storeDQ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DR() *comp.DR {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 576460752303423488 != 0 {
        return &storeDR[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DS() *comp.DS {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 1152921504606846976 != 0 {
        return &storeDS[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DT() *comp.DT {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 2305843009213693952 != 0 {
        return &storeDT[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DU() *comp.DU {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & 4611686018427387904 != 0 {
        return &storeDU[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DV() *comp.DV {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[1] & -9223372036854775808 != 0 {
        return &storeDV[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DW() *comp.DW {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 1 != 0 {
        return &storeDW[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DX() *comp.DX {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 2 != 0 {
        return &storeDX[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DY() *comp.DY {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 4 != 0 {
        return &storeDY[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) DZ() *comp.DZ {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 8 != 0 {
        return &storeDZ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EA() *comp.EA {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 16 != 0 {
        return &storeEA[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EB() *comp.EB {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 32 != 0 {
        return &storeEB[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EC() *comp.EC {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 64 != 0 {
        return &storeEC[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) ED() *comp.ED {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 128 != 0 {
        return &storeED[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EE() *comp.EE {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 256 != 0 {
        return &storeEE[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EF() *comp.EF {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 512 != 0 {
        return &storeEF[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EG() *comp.EG {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 1024 != 0 {
        return &storeEG[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EH() *comp.EH {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 2048 != 0 {
        return &storeEH[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EI() *comp.EI {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 4096 != 0 {
        return &storeEI[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EJ() *comp.EJ {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 8192 != 0 {
        return &storeEJ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EK() *comp.EK {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 16384 != 0 {
        return &storeEK[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EL() *comp.EL {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 32768 != 0 {
        return &storeEL[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EM() *comp.EM {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 65536 != 0 {
        return &storeEM[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EN() *comp.EN {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 131072 != 0 {
        return &storeEN[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EO() *comp.EO {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 262144 != 0 {
        return &storeEO[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EP() *comp.EP {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 524288 != 0 {
        return &storeEP[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) EQ() *comp.EQ {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 1048576 != 0 {
        return &storeEQ[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) ER() *comp.ER {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 2097152 != 0 {
        return &storeER[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) ES() *comp.ES {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 4194304 != 0 {
        return &storeES[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}

func (e Entity) ET() *comp.ET {
    if !e.Alive() {
        return nil
    }
    if entities[e.id >> entityPageBits][e.id % entityPageSize].components[2] & 8388608 != 0 {
        return &storeET[e.id >> entityPageBits][e.id % entityPageSize]
    }
    return nil
}



func (e Entity) DefaultPosition(def comp.Position) *comp.Position {
    if !e.HasPosition() {
        e.SetPosition(def)
    }

    return &storePosition[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultVelocity(def comp.Velocity) *comp.Velocity {
    if !e.HasVelocity() {
        e.SetVelocity(def)
    }

    return &storeVelocity[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultA(def comp.A) *comp.A {
    if !e.HasA() {
        e.SetA(def)
    }

    return &storeA[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultB(def comp.B) *comp.B {
    if !e.HasB() {
        e.SetB(def)
    }

    return &storeB[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultC(def comp.C) *comp.C {
    if !e.HasC() {
        e.SetC(def)
    }

    return &storeC[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultD(def comp.D) *comp.D {
    if !e.HasD() {
        e.SetD(def)
    }

    return &storeD[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultE(def comp.E) *comp.E {
    if !e.HasE() {
        e.SetE(def)
    }

    return &storeE[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultF(def comp.F) *comp.F {
    if !e.HasF() {
        e.SetF(def)
    }

    return &storeF[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultG(def comp.G) *comp.G {
    if !e.HasG() {
        e.SetG(def)
    }

    return &storeG[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultH(def comp.H) *comp.H {
    if !e.HasH() {
        e.SetH(def)
    }

    return &storeH[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultI(def comp.I) *comp.I {
    if !e.HasI() {
        e.SetI(def)
    }

    return &storeI[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultJ(def comp.J) *comp.J {
    if !e.HasJ() {
        e.SetJ(def)
    }

    return &storeJ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultK(def comp.K) *comp.K {
    if !e.HasK() {
        e.SetK(def)
    }

    return &storeK[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultL(def comp.L) *comp.L {
    if !e.HasL() {
        e.SetL(def)
    }

    return &storeL[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultM(def comp.M) *comp.M {
    if !e.HasM() {
        e.SetM(def)
    }

    return &storeM[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultN(def comp.N) *comp.N {
    if !e.HasN() {
        e.SetN(def)
    }

    return &storeN[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultO(def comp.O) *comp.O {
    if !e.HasO() {
        e.SetO(def)
    }

    return &storeO[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultP(def comp.P) *comp.P {
    if !e.HasP() {
        e.SetP(def)
    }

    return &storeP[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultQ(def comp.Q) *comp.Q {
    if !e.HasQ() {
        e.SetQ(def)
    }

    return &storeQ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultR(def comp.R) *comp.R {
    if !e.HasR() {
        e.SetR(def)
    }

    return &storeR[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultS(def comp.S) *comp.S {
    if !e.HasS() {
        e.SetS(def)
    }

    return &storeS[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultT(def comp.T) *comp.T {
    if !e.HasT() {
        e.SetT(def)
    }

    return &storeT[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultU(def comp.U) *comp.U {
    if !e.HasU() {
        e.SetU(def)
    }

    return &storeU[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultV(def comp.V) *comp.V {
    if !e.HasV() {
        e.SetV(def)
    }

    return &storeV[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultW(def comp.W) *comp.W {
    if !e.HasW() {
        e.SetW(def)
    }

    return &storeW[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultX(def comp.X) *comp.X {
    if !e.HasX() {
        e.SetX(def)
    }

    return &storeX[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultY(def comp.Y) *comp.Y {
    if !e.HasY() {
        e.SetY(def)
    }

    return &storeY[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultZ(def comp.Z) *comp.Z {
    if !e.HasZ() {
        e.SetZ(def)
    }

    return &storeZ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAA(def comp.AA) *comp.AA {
    if !e.HasAA() {
        e.SetAA(def)
    }

    return &storeAA[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAB(def comp.AB) *comp.AB {
    if !e.HasAB() {
        e.SetAB(def)
    }

    return &storeAB[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAC(def comp.AC) *comp.AC {
    if !e.HasAC() {
        e.SetAC(def)
    }

    return &storeAC[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAD(def comp.AD) *comp.AD {
    if !e.HasAD() {
        e.SetAD(def)
    }

    return &storeAD[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAE(def comp.AE) *comp.AE {
    if !e.HasAE() {
        e.SetAE(def)
    }

    return &storeAE[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAF(def comp.AF) *comp.AF {
    if !e.HasAF() {
        e.SetAF(def)
    }

    return &storeAF[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAG(def comp.AG) *comp.AG {
    if !e.HasAG() {
        e.SetAG(def)
    }

    return &storeAG[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAH(def comp.AH) *comp.AH {
    if !e.HasAH() {
        e.SetAH(def)
    }

    return &storeAH[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAI(def comp.AI) *comp.AI {
    if !e.HasAI() {
        e.SetAI(def)
    }

    return &storeAI[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAJ(def comp.AJ) *comp.AJ {
    if !e.HasAJ() {
        e.SetAJ(def)
    }

    return &storeAJ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAK(def comp.AK) *comp.AK {
    if !e.HasAK() {
        e.SetAK(def)
    }

    return &storeAK[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAL(def comp.AL) *comp.AL {
    if !e.HasAL() {
        e.SetAL(def)
    }

    return &storeAL[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAM(def comp.AM) *comp.AM {
    if !e.HasAM() {
        e.SetAM(def)
    }

    return &storeAM[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAN(def comp.AN) *comp.AN {
    if !e.HasAN() {
        e.SetAN(def)
    }

    return &storeAN[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAO(def comp.AO) *comp.AO {
    if !e.HasAO() {
        e.SetAO(def)
    }

    return &storeAO[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAP(def comp.AP) *comp.AP {
    if !e.HasAP() {
        e.SetAP(def)
    }

    return &storeAP[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAQ(def comp.AQ) *comp.AQ {
    if !e.HasAQ() {
        e.SetAQ(def)
    }

    return &storeAQ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAR(def comp.AR) *comp.AR {
    if !e.HasAR() {
        e.SetAR(def)
    }

    return &storeAR[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAS(def comp.AS) *comp.AS {
    if !e.HasAS() {
        e.SetAS(def)
    }

    return &storeAS[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAT(def comp.AT) *comp.AT {
    if !e.HasAT() {
        e.SetAT(def)
    }

    return &storeAT[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAU(def comp.AU) *comp.AU {
    if !e.HasAU() {
        e.SetAU(def)
    }

    return &storeAU[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAV(def comp.AV) *comp.AV {
    if !e.HasAV() {
        e.SetAV(def)
    }

    return &storeAV[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAW(def comp.AW) *comp.AW {
    if !e.HasAW() {
        e.SetAW(def)
    }

    return &storeAW[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAX(def comp.AX) *comp.AX {
    if !e.HasAX() {
        e.SetAX(def)
    }

    return &storeAX[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAY(def comp.AY) *comp.AY {
    if !e.HasAY() {
        e.SetAY(def)
    }

    return &storeAY[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultAZ(def comp.AZ) *comp.AZ {
    if !e.HasAZ() {
        e.SetAZ(def)
    }

    return &storeAZ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBA(def comp.BA) *comp.BA {
    if !e.HasBA() {
        e.SetBA(def)
    }

    return &storeBA[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBB(def comp.BB) *comp.BB {
    if !e.HasBB() {
        e.SetBB(def)
    }

    return &storeBB[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBC(def comp.BC) *comp.BC {
    if !e.HasBC() {
        e.SetBC(def)
    }

    return &storeBC[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBD(def comp.BD) *comp.BD {
    if !e.HasBD() {
        e.SetBD(def)
    }

    return &storeBD[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBE(def comp.BE) *comp.BE {
    if !e.HasBE() {
        e.SetBE(def)
    }

    return &storeBE[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBF(def comp.BF) *comp.BF {
    if !e.HasBF() {
        e.SetBF(def)
    }

    return &storeBF[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBG(def comp.BG) *comp.BG {
    if !e.HasBG() {
        e.SetBG(def)
    }

    return &storeBG[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBH(def comp.BH) *comp.BH {
    if !e.HasBH() {
        e.SetBH(def)
    }

    return &storeBH[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBI(def comp.BI) *comp.BI {
    if !e.HasBI() {
        e.SetBI(def)
    }

    return &storeBI[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBJ(def comp.BJ) *comp.BJ {
    if !e.HasBJ() {
        e.SetBJ(def)
    }

    return &storeBJ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBK(def comp.BK) *comp.BK {
    if !e.HasBK() {
        e.SetBK(def)
    }

    return &storeBK[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBL(def comp.BL) *comp.BL {
    if !e.HasBL() {
        e.SetBL(def)
    }

    return &storeBL[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBM(def comp.BM) *comp.BM {
    if !e.HasBM() {
        e.SetBM(def)
    }

    return &storeBM[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBN(def comp.BN) *comp.BN {
    if !e.HasBN() {
        e.SetBN(def)
    }

    return &storeBN[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBO(def comp.BO) *comp.BO {
    if !e.HasBO() {
        e.SetBO(def)
    }

    return &storeBO[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBP(def comp.BP) *comp.BP {
    if !e.HasBP() {
        e.SetBP(def)
    }

    return &storeBP[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBQ(def comp.BQ) *comp.BQ {
    if !e.HasBQ() {
        e.SetBQ(def)
    }

    return &storeBQ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBR(def comp.BR) *comp.BR {
    if !e.HasBR() {
        e.SetBR(def)
    }

    return &storeBR[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBS(def comp.BS) *comp.BS {
    if !e.HasBS() {
        e.SetBS(def)
    }

    return &storeBS[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBT(def comp.BT) *comp.BT {
    if !e.HasBT() {
        e.SetBT(def)
    }

    return &storeBT[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBU(def comp.BU) *comp.BU {
    if !e.HasBU() {
        e.SetBU(def)
    }

    return &storeBU[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBV(def comp.BV) *comp.BV {
    if !e.HasBV() {
        e.SetBV(def)
    }

    return &storeBV[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBW(def comp.BW) *comp.BW {
    if !e.HasBW() {
        e.SetBW(def)
    }

    return &storeBW[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBX(def comp.BX) *comp.BX {
    if !e.HasBX() {
        e.SetBX(def)
    }

    return &storeBX[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBY(def comp.BY) *comp.BY {
    if !e.HasBY() {
        e.SetBY(def)
    }

    return &storeBY[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultBZ(def comp.BZ) *comp.BZ {
    if !e.HasBZ() {
        e.SetBZ(def)
    }

    return &storeBZ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCA(def comp.CA) *comp.CA {
    if !e.HasCA() {
        e.SetCA(def)
    }

    return &storeCA[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCB(def comp.CB) *comp.CB {
    if !e.HasCB() {
        e.SetCB(def)
    }

    return &storeCB[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCC(def comp.CC) *comp.CC {
    if !e.HasCC() {
        e.SetCC(def)
    }

    return &storeCC[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCD(def comp.CD) *comp.CD {
    if !e.HasCD() {
        e.SetCD(def)
    }

    return &storeCD[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCE(def comp.CE) *comp.CE {
    if !e.HasCE() {
        e.SetCE(def)
    }

    return &storeCE[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCF(def comp.CF) *comp.CF {
    if !e.HasCF() {
        e.SetCF(def)
    }

    return &storeCF[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCG(def comp.CG) *comp.CG {
    if !e.HasCG() {
        e.SetCG(def)
    }

    return &storeCG[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCH(def comp.CH) *comp.CH {
    if !e.HasCH() {
        e.SetCH(def)
    }

    return &storeCH[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCI(def comp.CI) *comp.CI {
    if !e.HasCI() {
        e.SetCI(def)
    }

    return &storeCI[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCJ(def comp.CJ) *comp.CJ {
    if !e.HasCJ() {
        e.SetCJ(def)
    }

    return &storeCJ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCK(def comp.CK) *comp.CK {
    if !e.HasCK() {
        e.SetCK(def)
    }

    return &storeCK[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCL(def comp.CL) *comp.CL {
    if !e.HasCL() {
        e.SetCL(def)
    }

    return &storeCL[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCM(def comp.CM) *comp.CM {
    if !e.HasCM() {
        e.SetCM(def)
    }

    return &storeCM[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCN(def comp.CN) *comp.CN {
    if !e.HasCN() {
        e.SetCN(def)
    }

    return &storeCN[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCO(def comp.CO) *comp.CO {
    if !e.HasCO() {
        e.SetCO(def)
    }

    return &storeCO[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCP(def comp.CP) *comp.CP {
    if !e.HasCP() {
        e.SetCP(def)
    }

    return &storeCP[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCQ(def comp.CQ) *comp.CQ {
    if !e.HasCQ() {
        e.SetCQ(def)
    }

    return &storeCQ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCR(def comp.CR) *comp.CR {
    if !e.HasCR() {
        e.SetCR(def)
    }

    return &storeCR[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCS(def comp.CS) *comp.CS {
    if !e.HasCS() {
        e.SetCS(def)
    }

    return &storeCS[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCT(def comp.CT) *comp.CT {
    if !e.HasCT() {
        e.SetCT(def)
    }

    return &storeCT[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCU(def comp.CU) *comp.CU {
    if !e.HasCU() {
        e.SetCU(def)
    }

    return &storeCU[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCV(def comp.CV) *comp.CV {
    if !e.HasCV() {
        e.SetCV(def)
    }

    return &storeCV[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCW(def comp.CW) *comp.CW {
    if !e.HasCW() {
        e.SetCW(def)
    }

    return &storeCW[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCX(def comp.CX) *comp.CX {
    if !e.HasCX() {
        e.SetCX(def)
    }

    return &storeCX[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCY(def comp.CY) *comp.CY {
    if !e.HasCY() {
        e.SetCY(def)
    }

    return &storeCY[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultCZ(def comp.CZ) *comp.CZ {
    if !e.HasCZ() {
        e.SetCZ(def)
    }

    return &storeCZ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDA(def comp.DA) *comp.DA {
    if !e.HasDA() {
        e.SetDA(def)
    }

    return &storeDA[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDB(def comp.DB) *comp.DB {
    if !e.HasDB() {
        e.SetDB(def)
    }

    return &storeDB[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDC(def comp.DC) *comp.DC {
    if !e.HasDC() {
        e.SetDC(def)
    }

    return &storeDC[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDD(def comp.DD) *comp.DD {
    if !e.HasDD() {
        e.SetDD(def)
    }

    return &storeDD[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDE(def comp.DE) *comp.DE {
    if !e.HasDE() {
        e.SetDE(def)
    }

    return &storeDE[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDF(def comp.DF) *comp.DF {
    if !e.HasDF() {
        e.SetDF(def)
    }

    return &storeDF[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDG(def comp.DG) *comp.DG {
    if !e.HasDG() {
        e.SetDG(def)
    }

    return &storeDG[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDH(def comp.DH) *comp.DH {
    if !e.HasDH() {
        e.SetDH(def)
    }

    return &storeDH[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDI(def comp.DI) *comp.DI {
    if !e.HasDI() {
        e.SetDI(def)
    }

    return &storeDI[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDJ(def comp.DJ) *comp.DJ {
    if !e.HasDJ() {
        e.SetDJ(def)
    }

    return &storeDJ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDK(def comp.DK) *comp.DK {
    if !e.HasDK() {
        e.SetDK(def)
    }

    return &storeDK[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDL(def comp.DL) *comp.DL {
    if !e.HasDL() {
        e.SetDL(def)
    }

    return &storeDL[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDM(def comp.DM) *comp.DM {
    if !e.HasDM() {
        e.SetDM(def)
    }

    return &storeDM[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDN(def comp.DN) *comp.DN {
    if !e.HasDN() {
        e.SetDN(def)
    }

    return &storeDN[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDO(def comp.DO) *comp.DO {
    if !e.HasDO() {
        e.SetDO(def)
    }

    return &storeDO[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDP(def comp.DP) *comp.DP {
    if !e.HasDP() {
        e.SetDP(def)
    }

    return &storeDP[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDQ(def comp.DQ) *comp.DQ {
    if !e.HasDQ() {
        e.SetDQ(def)
    }

    return &storeDQ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDR(def comp.DR) *comp.DR {
    if !e.HasDR() {
        e.SetDR(def)
    }

    return &storeDR[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDS(def comp.DS) *comp.DS {
    if !e.HasDS() {
        e.SetDS(def)
    }

    return &storeDS[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDT(def comp.DT) *comp.DT {
    if !e.HasDT() {
        e.SetDT(def)
    }

    return &storeDT[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDU(def comp.DU) *comp.DU {
    if !e.HasDU() {
        e.SetDU(def)
    }

    return &storeDU[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDV(def comp.DV) *comp.DV {
    if !e.HasDV() {
        e.SetDV(def)
    }

    return &storeDV[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDW(def comp.DW) *comp.DW {
    if !e.HasDW() {
        e.SetDW(def)
    }

    return &storeDW[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDX(def comp.DX) *comp.DX {
    if !e.HasDX() {
        e.SetDX(def)
    }

    return &storeDX[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDY(def comp.DY) *comp.DY {
    if !e.HasDY() {
        e.SetDY(def)
    }

    return &storeDY[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultDZ(def comp.DZ) *comp.DZ {
    if !e.HasDZ() {
        e.SetDZ(def)
    }

    return &storeDZ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEA(def comp.EA) *comp.EA {
    if !e.HasEA() {
        e.SetEA(def)
    }

    return &storeEA[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEB(def comp.EB) *comp.EB {
    if !e.HasEB() {
        e.SetEB(def)
    }

    return &storeEB[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEC(def comp.EC) *comp.EC {
    if !e.HasEC() {
        e.SetEC(def)
    }

    return &storeEC[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultED(def comp.ED) *comp.ED {
    if !e.HasED() {
        e.SetED(def)
    }

    return &storeED[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEE(def comp.EE) *comp.EE {
    if !e.HasEE() {
        e.SetEE(def)
    }

    return &storeEE[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEF(def comp.EF) *comp.EF {
    if !e.HasEF() {
        e.SetEF(def)
    }

    return &storeEF[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEG(def comp.EG) *comp.EG {
    if !e.HasEG() {
        e.SetEG(def)
    }

    return &storeEG[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEH(def comp.EH) *comp.EH {
    if !e.HasEH() {
        e.SetEH(def)
    }

    return &storeEH[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEI(def comp.EI) *comp.EI {
    if !e.HasEI() {
        e.SetEI(def)
    }

    return &storeEI[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEJ(def comp.EJ) *comp.EJ {
    if !e.HasEJ() {
        e.SetEJ(def)
    }

    return &storeEJ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEK(def comp.EK) *comp.EK {
    if !e.HasEK() {
        e.SetEK(def)
    }

    return &storeEK[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEL(def comp.EL) *comp.EL {
    if !e.HasEL() {
        e.SetEL(def)
    }

    return &storeEL[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEM(def comp.EM) *comp.EM {
    if !e.HasEM() {
        e.SetEM(def)
    }

    return &storeEM[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEN(def comp.EN) *comp.EN {
    if !e.HasEN() {
        e.SetEN(def)
    }

    return &storeEN[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEO(def comp.EO) *comp.EO {
    if !e.HasEO() {
        e.SetEO(def)
    }

    return &storeEO[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEP(def comp.EP) *comp.EP {
    if !e.HasEP() {
        e.SetEP(def)
    }

    return &storeEP[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultEQ(def comp.EQ) *comp.EQ {
    if !e.HasEQ() {
        e.SetEQ(def)
    }

    return &storeEQ[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultER(def comp.ER) *comp.ER {
    if !e.HasER() {
        e.SetER(def)
    }

    return &storeER[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultES(def comp.ES) *comp.ES {
    if !e.HasES() {
        e.SetES(def)
    }

    return &storeES[e.id >> entityPageBits][e.id % entityPageSize]
}

func (e Entity) DefaultET(def comp.ET) *comp.ET {
    if !e.HasET() {
        e.SetET(def)
    }

    return &storeET[e.id >> entityPageBits][e.id % entityPageSize]
}


// Components returns all components on an entity, but it isn't fast - don't use for non-debug purposes!
func (e Entity) Components() []interface{} {
    comps := make([]interface{}, 0)
    
    if e.HasPosition() {
        comps = append(comps, *e.Position())
    }
    
    if e.HasVelocity() {
        comps = append(comps, *e.Velocity())
    }
    
    if e.HasA() {
        comps = append(comps, *e.A())
    }
    
    if e.HasB() {
        comps = append(comps, *e.B())
    }
    
    if e.HasC() {
        comps = append(comps, *e.C())
    }
    
    if e.HasD() {
        comps = append(comps, *e.D())
    }
    
    if e.HasE() {
        comps = append(comps, *e.E())
    }
    
    if e.HasF() {
        comps = append(comps, *e.F())
    }
    
    if e.HasG() {
        comps = append(comps, *e.G())
    }
    
    if e.HasH() {
        comps = append(comps, *e.H())
    }
    
    if e.HasI() {
        comps = append(comps, *e.I())
    }
    
    if e.HasJ() {
        comps = append(comps, *e.J())
    }
    
    if e.HasK() {
        comps = append(comps, *e.K())
    }
    
    if e.HasL() {
        comps = append(comps, *e.L())
    }
    
    if e.HasM() {
        comps = append(comps, *e.M())
    }
    
    if e.HasN() {
        comps = append(comps, *e.N())
    }
    
    if e.HasO() {
        comps = append(comps, *e.O())
    }
    
    if e.HasP() {
        comps = append(comps, *e.P())
    }
    
    if e.HasQ() {
        comps = append(comps, *e.Q())
    }
    
    if e.HasR() {
        comps = append(comps, *e.R())
    }
    
    if e.HasS() {
        comps = append(comps, *e.S())
    }
    
    if e.HasT() {
        comps = append(comps, *e.T())
    }
    
    if e.HasU() {
        comps = append(comps, *e.U())
    }
    
    if e.HasV() {
        comps = append(comps, *e.V())
    }
    
    if e.HasW() {
        comps = append(comps, *e.W())
    }
    
    if e.HasX() {
        comps = append(comps, *e.X())
    }
    
    if e.HasY() {
        comps = append(comps, *e.Y())
    }
    
    if e.HasZ() {
        comps = append(comps, *e.Z())
    }
    
    if e.HasAA() {
        comps = append(comps, *e.AA())
    }
    
    if e.HasAB() {
        comps = append(comps, *e.AB())
    }
    
    if e.HasAC() {
        comps = append(comps, *e.AC())
    }
    
    if e.HasAD() {
        comps = append(comps, *e.AD())
    }
    
    if e.HasAE() {
        comps = append(comps, *e.AE())
    }
    
    if e.HasAF() {
        comps = append(comps, *e.AF())
    }
    
    if e.HasAG() {
        comps = append(comps, *e.AG())
    }
    
    if e.HasAH() {
        comps = append(comps, *e.AH())
    }
    
    if e.HasAI() {
        comps = append(comps, *e.AI())
    }
    
    if e.HasAJ() {
        comps = append(comps, *e.AJ())
    }
    
    if e.HasAK() {
        comps = append(comps, *e.AK())
    }
    
    if e.HasAL() {
        comps = append(comps, *e.AL())
    }
    
    if e.HasAM() {
        comps = append(comps, *e.AM())
    }
    
    if e.HasAN() {
        comps = append(comps, *e.AN())
    }
    
    if e.HasAO() {
        comps = append(comps, *e.AO())
    }
    
    if e.HasAP() {
        comps = append(comps, *e.AP())
    }
    
    if e.HasAQ() {
        comps = append(comps, *e.AQ())
    }
    
    if e.HasAR() {
        comps = append(comps, *e.AR())
    }
    
    if e.HasAS() {
        comps = append(comps, *e.AS())
    }
    
    if e.HasAT() {
        comps = append(comps, *e.AT())
    }
    
    if e.HasAU() {
        comps = append(comps, *e.AU())
    }
    
    if e.HasAV() {
        comps = append(comps, *e.AV())
    }
    
    if e.HasAW() {
        comps = append(comps, *e.AW())
    }
    
    if e.HasAX() {
        comps = append(comps, *e.AX())
    }
    
    if e.HasAY() {
        comps = append(comps, *e.AY())
    }
    
    if e.HasAZ() {
        comps = append(comps, *e.AZ())
    }
    
    if e.HasBA() {
        comps = append(comps, *e.BA())
    }
    
    if e.HasBB() {
        comps = append(comps, *e.BB())
    }
    
    if e.HasBC() {
        comps = append(comps, *e.BC())
    }
    
    if e.HasBD() {
        comps = append(comps, *e.BD())
    }
    
    if e.HasBE() {
        comps = append(comps, *e.BE())
    }
    
    if e.HasBF() {
        comps = append(comps, *e.BF())
    }
    
    if e.HasBG() {
        comps = append(comps, *e.BG())
    }
    
    if e.HasBH() {
        comps = append(comps, *e.BH())
    }
    
    if e.HasBI() {
        comps = append(comps, *e.BI())
    }
    
    if e.HasBJ() {
        comps = append(comps, *e.BJ())
    }
    
    if e.HasBK() {
        comps = append(comps, *e.BK())
    }
    
    if e.HasBL() {
        comps = append(comps, *e.BL())
    }
    
    if e.HasBM() {
        comps = append(comps, *e.BM())
    }
    
    if e.HasBN() {
        comps = append(comps, *e.BN())
    }
    
    if e.HasBO() {
        comps = append(comps, *e.BO())
    }
    
    if e.HasBP() {
        comps = append(comps, *e.BP())
    }
    
    if e.HasBQ() {
        comps = append(comps, *e.BQ())
    }
    
    if e.HasBR() {
        comps = append(comps, *e.BR())
    }
    
    if e.HasBS() {
        comps = append(comps, *e.BS())
    }
    
    if e.HasBT() {
        comps = append(comps, *e.BT())
    }
    
    if e.HasBU() {
        comps = append(comps, *e.BU())
    }
    
    if e.HasBV() {
        comps = append(comps, *e.BV())
    }
    
    if e.HasBW() {
        comps = append(comps, *e.BW())
    }
    
    if e.HasBX() {
        comps = append(comps, *e.BX())
    }
    
    if e.HasBY() {
        comps = append(comps, *e.BY())
    }
    
    if e.HasBZ() {
        comps = append(comps, *e.BZ())
    }
    
    if e.HasCA() {
        comps = append(comps, *e.CA())
    }
    
    if e.HasCB() {
        comps = append(comps, *e.CB())
    }
    
    if e.HasCC() {
        comps = append(comps, *e.CC())
    }
    
    if e.HasCD() {
        comps = append(comps, *e.CD())
    }
    
    if e.HasCE() {
        comps = append(comps, *e.CE())
    }
    
    if e.HasCF() {
        comps = append(comps, *e.CF())
    }
    
    if e.HasCG() {
        comps = append(comps, *e.CG())
    }
    
    if e.HasCH() {
        comps = append(comps, *e.CH())
    }
    
    if e.HasCI() {
        comps = append(comps, *e.CI())
    }
    
    if e.HasCJ() {
        comps = append(comps, *e.CJ())
    }
    
    if e.HasCK() {
        comps = append(comps, *e.CK())
    }
    
    if e.HasCL() {
        comps = append(comps, *e.CL())
    }
    
    if e.HasCM() {
        comps = append(comps, *e.CM())
    }
    
    if e.HasCN() {
        comps = append(comps, *e.CN())
    }
    
    if e.HasCO() {
        comps = append(comps, *e.CO())
    }
    
    if e.HasCP() {
        comps = append(comps, *e.CP())
    }
    
    if e.HasCQ() {
        comps = append(comps, *e.CQ())
    }
    
    if e.HasCR() {
        comps = append(comps, *e.CR())
    }
    
    if e.HasCS() {
        comps = append(comps, *e.CS())
    }
    
    if e.HasCT() {
        comps = append(comps, *e.CT())
    }
    
    if e.HasCU() {
        comps = append(comps, *e.CU())
    }
    
    if e.HasCV() {
        comps = append(comps, *e.CV())
    }
    
    if e.HasCW() {
        comps = append(comps, *e.CW())
    }
    
    if e.HasCX() {
        comps = append(comps, *e.CX())
    }
    
    if e.HasCY() {
        comps = append(comps, *e.CY())
    }
    
    if e.HasCZ() {
        comps = append(comps, *e.CZ())
    }
    
    if e.HasDA() {
        comps = append(comps, *e.DA())
    }
    
    if e.HasDB() {
        comps = append(comps, *e.DB())
    }
    
    if e.HasDC() {
        comps = append(comps, *e.DC())
    }
    
    if e.HasDD() {
        comps = append(comps, *e.DD())
    }
    
    if e.HasDE() {
        comps = append(comps, *e.DE())
    }
    
    if e.HasDF() {
        comps = append(comps, *e.DF())
    }
    
    if e.HasDG() {
        comps = append(comps, *e.DG())
    }
    
    if e.HasDH() {
        comps = append(comps, *e.DH())
    }
    
    if e.HasDI() {
        comps = append(comps, *e.DI())
    }
    
    if e.HasDJ() {
        comps = append(comps, *e.DJ())
    }
    
    if e.HasDK() {
        comps = append(comps, *e.DK())
    }
    
    if e.HasDL() {
        comps = append(comps, *e.DL())
    }
    
    if e.HasDM() {
        comps = append(comps, *e.DM())
    }
    
    if e.HasDN() {
        comps = append(comps, *e.DN())
    }
    
    if e.HasDO() {
        comps = append(comps, *e.DO())
    }
    
    if e.HasDP() {
        comps = append(comps, *e.DP())
    }
    
    if e.HasDQ() {
        comps = append(comps, *e.DQ())
    }
    
    if e.HasDR() {
        comps = append(comps, *e.DR())
    }
    
    if e.HasDS() {
        comps = append(comps, *e.DS())
    }
    
    if e.HasDT() {
        comps = append(comps, *e.DT())
    }
    
    if e.HasDU() {
        comps = append(comps, *e.DU())
    }
    
    if e.HasDV() {
        comps = append(comps, *e.DV())
    }
    
    if e.HasDW() {
        comps = append(comps, *e.DW())
    }
    
    if e.HasDX() {
        comps = append(comps, *e.DX())
    }
    
    if e.HasDY() {
        comps = append(comps, *e.DY())
    }
    
    if e.HasDZ() {
        comps = append(comps, *e.DZ())
    }
    
    if e.HasEA() {
        comps = append(comps, *e.EA())
    }
    
    if e.HasEB() {
        comps = append(comps, *e.EB())
    }
    
    if e.HasEC() {
        comps = append(comps, *e.EC())
    }
    
    if e.HasED() {
        comps = append(comps, *e.ED())
    }
    
    if e.HasEE() {
        comps = append(comps, *e.EE())
    }
    
    if e.HasEF() {
        comps = append(comps, *e.EF())
    }
    
    if e.HasEG() {
        comps = append(comps, *e.EG())
    }
    
    if e.HasEH() {
        comps = append(comps, *e.EH())
    }
    
    if e.HasEI() {
        comps = append(comps, *e.EI())
    }
    
    if e.HasEJ() {
        comps = append(comps, *e.EJ())
    }
    
    if e.HasEK() {
        comps = append(comps, *e.EK())
    }
    
    if e.HasEL() {
        comps = append(comps, *e.EL())
    }
    
    if e.HasEM() {
        comps = append(comps, *e.EM())
    }
    
    if e.HasEN() {
        comps = append(comps, *e.EN())
    }
    
    if e.HasEO() {
        comps = append(comps, *e.EO())
    }
    
    if e.HasEP() {
        comps = append(comps, *e.EP())
    }
    
    if e.HasEQ() {
        comps = append(comps, *e.EQ())
    }
    
    if e.HasER() {
        comps = append(comps, *e.ER())
    }
    
    if e.HasES() {
        comps = append(comps, *e.ES())
    }
    
    if e.HasET() {
        comps = append(comps, *e.ET())
    }
    
    return comps
}