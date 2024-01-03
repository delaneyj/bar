package myecspkg

import comp "bar/components"
import "bar/myecspkg/entity"
import "unsafe"

var currEntities = 0
var entityCap = 0

var systems []System

func init() {
    newEntityPage()
}

var sortSpace []Entity

func newEntityPage() {
    newPage := make([]Entity, entityPageSize)
    entities = append(entities, newPage)
    sortSpace = make([]Entity, entityCap + entityPageSize)

    
    newPositionPage := make([]comp.Position, entityPageSize)
    storePosition = append(storePosition, newPositionPage)
    
    newVelocityPage := make([]comp.Velocity, entityPageSize)
    storeVelocity = append(storeVelocity, newVelocityPage)
    
    newAPage := make([]comp.A, entityPageSize)
    storeA = append(storeA, newAPage)
    
    newBPage := make([]comp.B, entityPageSize)
    storeB = append(storeB, newBPage)
    
    newCPage := make([]comp.C, entityPageSize)
    storeC = append(storeC, newCPage)
    
    newDPage := make([]comp.D, entityPageSize)
    storeD = append(storeD, newDPage)
    
    newEPage := make([]comp.E, entityPageSize)
    storeE = append(storeE, newEPage)
    
    newFPage := make([]comp.F, entityPageSize)
    storeF = append(storeF, newFPage)
    
    newGPage := make([]comp.G, entityPageSize)
    storeG = append(storeG, newGPage)
    
    newHPage := make([]comp.H, entityPageSize)
    storeH = append(storeH, newHPage)
    
    newIPage := make([]comp.I, entityPageSize)
    storeI = append(storeI, newIPage)
    
    newJPage := make([]comp.J, entityPageSize)
    storeJ = append(storeJ, newJPage)
    
    newKPage := make([]comp.K, entityPageSize)
    storeK = append(storeK, newKPage)
    
    newLPage := make([]comp.L, entityPageSize)
    storeL = append(storeL, newLPage)
    
    newMPage := make([]comp.M, entityPageSize)
    storeM = append(storeM, newMPage)
    
    newNPage := make([]comp.N, entityPageSize)
    storeN = append(storeN, newNPage)
    
    newOPage := make([]comp.O, entityPageSize)
    storeO = append(storeO, newOPage)
    
    newPPage := make([]comp.P, entityPageSize)
    storeP = append(storeP, newPPage)
    
    newQPage := make([]comp.Q, entityPageSize)
    storeQ = append(storeQ, newQPage)
    
    newRPage := make([]comp.R, entityPageSize)
    storeR = append(storeR, newRPage)
    
    newSPage := make([]comp.S, entityPageSize)
    storeS = append(storeS, newSPage)
    
    newTPage := make([]comp.T, entityPageSize)
    storeT = append(storeT, newTPage)
    
    newUPage := make([]comp.U, entityPageSize)
    storeU = append(storeU, newUPage)
    
    newVPage := make([]comp.V, entityPageSize)
    storeV = append(storeV, newVPage)
    
    newWPage := make([]comp.W, entityPageSize)
    storeW = append(storeW, newWPage)
    
    newXPage := make([]comp.X, entityPageSize)
    storeX = append(storeX, newXPage)
    
    newYPage := make([]comp.Y, entityPageSize)
    storeY = append(storeY, newYPage)
    
    newZPage := make([]comp.Z, entityPageSize)
    storeZ = append(storeZ, newZPage)
    
    newAAPage := make([]comp.AA, entityPageSize)
    storeAA = append(storeAA, newAAPage)
    
    newABPage := make([]comp.AB, entityPageSize)
    storeAB = append(storeAB, newABPage)
    
    newACPage := make([]comp.AC, entityPageSize)
    storeAC = append(storeAC, newACPage)
    
    newADPage := make([]comp.AD, entityPageSize)
    storeAD = append(storeAD, newADPage)
    
    newAEPage := make([]comp.AE, entityPageSize)
    storeAE = append(storeAE, newAEPage)
    
    newAFPage := make([]comp.AF, entityPageSize)
    storeAF = append(storeAF, newAFPage)
    
    newAGPage := make([]comp.AG, entityPageSize)
    storeAG = append(storeAG, newAGPage)
    
    newAHPage := make([]comp.AH, entityPageSize)
    storeAH = append(storeAH, newAHPage)
    
    newAIPage := make([]comp.AI, entityPageSize)
    storeAI = append(storeAI, newAIPage)
    
    newAJPage := make([]comp.AJ, entityPageSize)
    storeAJ = append(storeAJ, newAJPage)
    
    newAKPage := make([]comp.AK, entityPageSize)
    storeAK = append(storeAK, newAKPage)
    
    newALPage := make([]comp.AL, entityPageSize)
    storeAL = append(storeAL, newALPage)
    
    newAMPage := make([]comp.AM, entityPageSize)
    storeAM = append(storeAM, newAMPage)
    
    newANPage := make([]comp.AN, entityPageSize)
    storeAN = append(storeAN, newANPage)
    
    newAOPage := make([]comp.AO, entityPageSize)
    storeAO = append(storeAO, newAOPage)
    
    newAPPage := make([]comp.AP, entityPageSize)
    storeAP = append(storeAP, newAPPage)
    
    newAQPage := make([]comp.AQ, entityPageSize)
    storeAQ = append(storeAQ, newAQPage)
    
    newARPage := make([]comp.AR, entityPageSize)
    storeAR = append(storeAR, newARPage)
    
    newASPage := make([]comp.AS, entityPageSize)
    storeAS = append(storeAS, newASPage)
    
    newATPage := make([]comp.AT, entityPageSize)
    storeAT = append(storeAT, newATPage)
    
    newAUPage := make([]comp.AU, entityPageSize)
    storeAU = append(storeAU, newAUPage)
    
    newAVPage := make([]comp.AV, entityPageSize)
    storeAV = append(storeAV, newAVPage)
    
    newAWPage := make([]comp.AW, entityPageSize)
    storeAW = append(storeAW, newAWPage)
    
    newAXPage := make([]comp.AX, entityPageSize)
    storeAX = append(storeAX, newAXPage)
    
    newAYPage := make([]comp.AY, entityPageSize)
    storeAY = append(storeAY, newAYPage)
    
    newAZPage := make([]comp.AZ, entityPageSize)
    storeAZ = append(storeAZ, newAZPage)
    
    newBAPage := make([]comp.BA, entityPageSize)
    storeBA = append(storeBA, newBAPage)
    
    newBBPage := make([]comp.BB, entityPageSize)
    storeBB = append(storeBB, newBBPage)
    
    newBCPage := make([]comp.BC, entityPageSize)
    storeBC = append(storeBC, newBCPage)
    
    newBDPage := make([]comp.BD, entityPageSize)
    storeBD = append(storeBD, newBDPage)
    
    newBEPage := make([]comp.BE, entityPageSize)
    storeBE = append(storeBE, newBEPage)
    
    newBFPage := make([]comp.BF, entityPageSize)
    storeBF = append(storeBF, newBFPage)
    
    newBGPage := make([]comp.BG, entityPageSize)
    storeBG = append(storeBG, newBGPage)
    
    newBHPage := make([]comp.BH, entityPageSize)
    storeBH = append(storeBH, newBHPage)
    
    newBIPage := make([]comp.BI, entityPageSize)
    storeBI = append(storeBI, newBIPage)
    
    newBJPage := make([]comp.BJ, entityPageSize)
    storeBJ = append(storeBJ, newBJPage)
    
    newBKPage := make([]comp.BK, entityPageSize)
    storeBK = append(storeBK, newBKPage)
    
    newBLPage := make([]comp.BL, entityPageSize)
    storeBL = append(storeBL, newBLPage)
    
    newBMPage := make([]comp.BM, entityPageSize)
    storeBM = append(storeBM, newBMPage)
    
    newBNPage := make([]comp.BN, entityPageSize)
    storeBN = append(storeBN, newBNPage)
    
    newBOPage := make([]comp.BO, entityPageSize)
    storeBO = append(storeBO, newBOPage)
    
    newBPPage := make([]comp.BP, entityPageSize)
    storeBP = append(storeBP, newBPPage)
    
    newBQPage := make([]comp.BQ, entityPageSize)
    storeBQ = append(storeBQ, newBQPage)
    
    newBRPage := make([]comp.BR, entityPageSize)
    storeBR = append(storeBR, newBRPage)
    
    newBSPage := make([]comp.BS, entityPageSize)
    storeBS = append(storeBS, newBSPage)
    
    newBTPage := make([]comp.BT, entityPageSize)
    storeBT = append(storeBT, newBTPage)
    
    newBUPage := make([]comp.BU, entityPageSize)
    storeBU = append(storeBU, newBUPage)
    
    newBVPage := make([]comp.BV, entityPageSize)
    storeBV = append(storeBV, newBVPage)
    
    newBWPage := make([]comp.BW, entityPageSize)
    storeBW = append(storeBW, newBWPage)
    
    newBXPage := make([]comp.BX, entityPageSize)
    storeBX = append(storeBX, newBXPage)
    
    newBYPage := make([]comp.BY, entityPageSize)
    storeBY = append(storeBY, newBYPage)
    
    newBZPage := make([]comp.BZ, entityPageSize)
    storeBZ = append(storeBZ, newBZPage)
    
    newCAPage := make([]comp.CA, entityPageSize)
    storeCA = append(storeCA, newCAPage)
    
    newCBPage := make([]comp.CB, entityPageSize)
    storeCB = append(storeCB, newCBPage)
    
    newCCPage := make([]comp.CC, entityPageSize)
    storeCC = append(storeCC, newCCPage)
    
    newCDPage := make([]comp.CD, entityPageSize)
    storeCD = append(storeCD, newCDPage)
    
    newCEPage := make([]comp.CE, entityPageSize)
    storeCE = append(storeCE, newCEPage)
    
    newCFPage := make([]comp.CF, entityPageSize)
    storeCF = append(storeCF, newCFPage)
    
    newCGPage := make([]comp.CG, entityPageSize)
    storeCG = append(storeCG, newCGPage)
    
    newCHPage := make([]comp.CH, entityPageSize)
    storeCH = append(storeCH, newCHPage)
    
    newCIPage := make([]comp.CI, entityPageSize)
    storeCI = append(storeCI, newCIPage)
    
    newCJPage := make([]comp.CJ, entityPageSize)
    storeCJ = append(storeCJ, newCJPage)
    
    newCKPage := make([]comp.CK, entityPageSize)
    storeCK = append(storeCK, newCKPage)
    
    newCLPage := make([]comp.CL, entityPageSize)
    storeCL = append(storeCL, newCLPage)
    
    newCMPage := make([]comp.CM, entityPageSize)
    storeCM = append(storeCM, newCMPage)
    
    newCNPage := make([]comp.CN, entityPageSize)
    storeCN = append(storeCN, newCNPage)
    
    newCOPage := make([]comp.CO, entityPageSize)
    storeCO = append(storeCO, newCOPage)
    
    newCPPage := make([]comp.CP, entityPageSize)
    storeCP = append(storeCP, newCPPage)
    
    newCQPage := make([]comp.CQ, entityPageSize)
    storeCQ = append(storeCQ, newCQPage)
    
    newCRPage := make([]comp.CR, entityPageSize)
    storeCR = append(storeCR, newCRPage)
    
    newCSPage := make([]comp.CS, entityPageSize)
    storeCS = append(storeCS, newCSPage)
    
    newCTPage := make([]comp.CT, entityPageSize)
    storeCT = append(storeCT, newCTPage)
    
    newCUPage := make([]comp.CU, entityPageSize)
    storeCU = append(storeCU, newCUPage)
    
    newCVPage := make([]comp.CV, entityPageSize)
    storeCV = append(storeCV, newCVPage)
    
    newCWPage := make([]comp.CW, entityPageSize)
    storeCW = append(storeCW, newCWPage)
    
    newCXPage := make([]comp.CX, entityPageSize)
    storeCX = append(storeCX, newCXPage)
    
    newCYPage := make([]comp.CY, entityPageSize)
    storeCY = append(storeCY, newCYPage)
    
    newCZPage := make([]comp.CZ, entityPageSize)
    storeCZ = append(storeCZ, newCZPage)
    
    newDAPage := make([]comp.DA, entityPageSize)
    storeDA = append(storeDA, newDAPage)
    
    newDBPage := make([]comp.DB, entityPageSize)
    storeDB = append(storeDB, newDBPage)
    
    newDCPage := make([]comp.DC, entityPageSize)
    storeDC = append(storeDC, newDCPage)
    
    newDDPage := make([]comp.DD, entityPageSize)
    storeDD = append(storeDD, newDDPage)
    
    newDEPage := make([]comp.DE, entityPageSize)
    storeDE = append(storeDE, newDEPage)
    
    newDFPage := make([]comp.DF, entityPageSize)
    storeDF = append(storeDF, newDFPage)
    
    newDGPage := make([]comp.DG, entityPageSize)
    storeDG = append(storeDG, newDGPage)
    
    newDHPage := make([]comp.DH, entityPageSize)
    storeDH = append(storeDH, newDHPage)
    
    newDIPage := make([]comp.DI, entityPageSize)
    storeDI = append(storeDI, newDIPage)
    
    newDJPage := make([]comp.DJ, entityPageSize)
    storeDJ = append(storeDJ, newDJPage)
    
    newDKPage := make([]comp.DK, entityPageSize)
    storeDK = append(storeDK, newDKPage)
    
    newDLPage := make([]comp.DL, entityPageSize)
    storeDL = append(storeDL, newDLPage)
    
    newDMPage := make([]comp.DM, entityPageSize)
    storeDM = append(storeDM, newDMPage)
    
    newDNPage := make([]comp.DN, entityPageSize)
    storeDN = append(storeDN, newDNPage)
    
    newDOPage := make([]comp.DO, entityPageSize)
    storeDO = append(storeDO, newDOPage)
    
    newDPPage := make([]comp.DP, entityPageSize)
    storeDP = append(storeDP, newDPPage)
    
    newDQPage := make([]comp.DQ, entityPageSize)
    storeDQ = append(storeDQ, newDQPage)
    
    newDRPage := make([]comp.DR, entityPageSize)
    storeDR = append(storeDR, newDRPage)
    
    newDSPage := make([]comp.DS, entityPageSize)
    storeDS = append(storeDS, newDSPage)
    
    newDTPage := make([]comp.DT, entityPageSize)
    storeDT = append(storeDT, newDTPage)
    
    newDUPage := make([]comp.DU, entityPageSize)
    storeDU = append(storeDU, newDUPage)
    
    newDVPage := make([]comp.DV, entityPageSize)
    storeDV = append(storeDV, newDVPage)
    
    newDWPage := make([]comp.DW, entityPageSize)
    storeDW = append(storeDW, newDWPage)
    
    newDXPage := make([]comp.DX, entityPageSize)
    storeDX = append(storeDX, newDXPage)
    
    newDYPage := make([]comp.DY, entityPageSize)
    storeDY = append(storeDY, newDYPage)
    
    newDZPage := make([]comp.DZ, entityPageSize)
    storeDZ = append(storeDZ, newDZPage)
    
    newEAPage := make([]comp.EA, entityPageSize)
    storeEA = append(storeEA, newEAPage)
    
    newEBPage := make([]comp.EB, entityPageSize)
    storeEB = append(storeEB, newEBPage)
    
    newECPage := make([]comp.EC, entityPageSize)
    storeEC = append(storeEC, newECPage)
    
    newEDPage := make([]comp.ED, entityPageSize)
    storeED = append(storeED, newEDPage)
    
    newEEPage := make([]comp.EE, entityPageSize)
    storeEE = append(storeEE, newEEPage)
    
    newEFPage := make([]comp.EF, entityPageSize)
    storeEF = append(storeEF, newEFPage)
    
    newEGPage := make([]comp.EG, entityPageSize)
    storeEG = append(storeEG, newEGPage)
    
    newEHPage := make([]comp.EH, entityPageSize)
    storeEH = append(storeEH, newEHPage)
    
    newEIPage := make([]comp.EI, entityPageSize)
    storeEI = append(storeEI, newEIPage)
    
    newEJPage := make([]comp.EJ, entityPageSize)
    storeEJ = append(storeEJ, newEJPage)
    
    newEKPage := make([]comp.EK, entityPageSize)
    storeEK = append(storeEK, newEKPage)
    
    newELPage := make([]comp.EL, entityPageSize)
    storeEL = append(storeEL, newELPage)
    
    newEMPage := make([]comp.EM, entityPageSize)
    storeEM = append(storeEM, newEMPage)
    
    newENPage := make([]comp.EN, entityPageSize)
    storeEN = append(storeEN, newENPage)
    
    newEOPage := make([]comp.EO, entityPageSize)
    storeEO = append(storeEO, newEOPage)
    
    newEPPage := make([]comp.EP, entityPageSize)
    storeEP = append(storeEP, newEPPage)
    
    newEQPage := make([]comp.EQ, entityPageSize)
    storeEQ = append(storeEQ, newEQPage)
    
    newERPage := make([]comp.ER, entityPageSize)
    storeER = append(storeER, newERPage)
    
    newESPage := make([]comp.ES, entityPageSize)
    storeES = append(storeES, newESPage)
    
    newETPage := make([]comp.ET, entityPageSize)
    storeET = append(storeET, newETPage)
    
    entityCap += entityPageSize
}

func Reset() {
    entities = nil
    
    storePosition = nil
    
    storeVelocity = nil
    
    storeA = nil
    
    storeB = nil
    
    storeC = nil
    
    storeD = nil
    
    storeE = nil
    
    storeF = nil
    
    storeG = nil
    
    storeH = nil
    
    storeI = nil
    
    storeJ = nil
    
    storeK = nil
    
    storeL = nil
    
    storeM = nil
    
    storeN = nil
    
    storeO = nil
    
    storeP = nil
    
    storeQ = nil
    
    storeR = nil
    
    storeS = nil
    
    storeT = nil
    
    storeU = nil
    
    storeV = nil
    
    storeW = nil
    
    storeX = nil
    
    storeY = nil
    
    storeZ = nil
    
    storeAA = nil
    
    storeAB = nil
    
    storeAC = nil
    
    storeAD = nil
    
    storeAE = nil
    
    storeAF = nil
    
    storeAG = nil
    
    storeAH = nil
    
    storeAI = nil
    
    storeAJ = nil
    
    storeAK = nil
    
    storeAL = nil
    
    storeAM = nil
    
    storeAN = nil
    
    storeAO = nil
    
    storeAP = nil
    
    storeAQ = nil
    
    storeAR = nil
    
    storeAS = nil
    
    storeAT = nil
    
    storeAU = nil
    
    storeAV = nil
    
    storeAW = nil
    
    storeAX = nil
    
    storeAY = nil
    
    storeAZ = nil
    
    storeBA = nil
    
    storeBB = nil
    
    storeBC = nil
    
    storeBD = nil
    
    storeBE = nil
    
    storeBF = nil
    
    storeBG = nil
    
    storeBH = nil
    
    storeBI = nil
    
    storeBJ = nil
    
    storeBK = nil
    
    storeBL = nil
    
    storeBM = nil
    
    storeBN = nil
    
    storeBO = nil
    
    storeBP = nil
    
    storeBQ = nil
    
    storeBR = nil
    
    storeBS = nil
    
    storeBT = nil
    
    storeBU = nil
    
    storeBV = nil
    
    storeBW = nil
    
    storeBX = nil
    
    storeBY = nil
    
    storeBZ = nil
    
    storeCA = nil
    
    storeCB = nil
    
    storeCC = nil
    
    storeCD = nil
    
    storeCE = nil
    
    storeCF = nil
    
    storeCG = nil
    
    storeCH = nil
    
    storeCI = nil
    
    storeCJ = nil
    
    storeCK = nil
    
    storeCL = nil
    
    storeCM = nil
    
    storeCN = nil
    
    storeCO = nil
    
    storeCP = nil
    
    storeCQ = nil
    
    storeCR = nil
    
    storeCS = nil
    
    storeCT = nil
    
    storeCU = nil
    
    storeCV = nil
    
    storeCW = nil
    
    storeCX = nil
    
    storeCY = nil
    
    storeCZ = nil
    
    storeDA = nil
    
    storeDB = nil
    
    storeDC = nil
    
    storeDD = nil
    
    storeDE = nil
    
    storeDF = nil
    
    storeDG = nil
    
    storeDH = nil
    
    storeDI = nil
    
    storeDJ = nil
    
    storeDK = nil
    
    storeDL = nil
    
    storeDM = nil
    
    storeDN = nil
    
    storeDO = nil
    
    storeDP = nil
    
    storeDQ = nil
    
    storeDR = nil
    
    storeDS = nil
    
    storeDT = nil
    
    storeDU = nil
    
    storeDV = nil
    
    storeDW = nil
    
    storeDX = nil
    
    storeDY = nil
    
    storeDZ = nil
    
    storeEA = nil
    
    storeEB = nil
    
    storeEC = nil
    
    storeED = nil
    
    storeEE = nil
    
    storeEF = nil
    
    storeEG = nil
    
    storeEH = nil
    
    storeEI = nil
    
    storeEJ = nil
    
    storeEK = nil
    
    storeEL = nil
    
    storeEM = nil
    
    storeEN = nil
    
    storeEO = nil
    
    storeEP = nil
    
    storeEQ = nil
    
    storeER = nil
    
    storeES = nil
    
    storeET = nil
    
    currEntities = 0
    entityCap = 0

    newEntityPage()
}

func NewEntity() Entity {
    var retID EntityID
    if len(freeList) == 0 {
        if currEntities >= entityCap {
            newEntityPage()
        }

        retID = EntityID(currEntities)
        currEntities++
    } else {
        retID = freeList[len(freeList)-1]
        freeList = freeList[:len(freeList)-1]
    }

	entities[retID >> entityPageBits][retID % entityPageSize] = Entity{
	    id: retID,
	    generation: entities[retID >> entityPageBits][retID % entityPageSize].generation + 1}

	return entities[retID >> entityPageBits][retID % entityPageSize]
}

func Lookup(r entity.Ref) Entity {
    return *(*Entity)(unsafe.Pointer(&r))
}

func (e Entity) Is(other Entity) bool {
    return e.id == other.id && e.generation == other.generation
}

func (e Entity) Zero() bool {
    return e.id == 0 && e.generation == 0
}

func (e Entity) Ref() entity.Ref {
    return *(*entity.Ref)(unsafe.Pointer(&e))
}

func EntityCount() int {
    return currEntities
}