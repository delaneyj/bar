package myecspkg

import comp "bar/components"
import "fmt"
import "reflect"
import "slices"
import "sync"

var sortLock sync.Mutex

func SelectSorted(cmp func(a Entity, b Entity) int, selector interface{}) {
    sortLock.Lock()
    defer sortLock.Unlock()

    i := 0
    switch selector.(type) {
    
    case func(Entity, *comp.BW, ):
        Select(func(e Entity, arg0 *comp.BW, ) {
            sortSpace[i] = e
            i++
        })
    
    case func(Entity, *comp.Position, *comp.Velocity, ):
        Select(func(e Entity, arg0 *comp.Position, arg1 *comp.Velocity, ) {
            sortSpace[i] = e
            i++
        })
    
    default:
        panic(fmt.Sprintf("unknown selector function: run go generate: %s", reflect.TypeOf(selector).String()))
    }

    slices.SortFunc(sortSpace[:i], cmp)
    for j := 0; j < i; j++ {
        entity := sortSpace[j]
        switch fun := selector.(type) {
        
        case func(Entity, *comp.BW, ):
            fun(entity, &storeBW[entity.id >> entityPageBits][entity.id % entityPageSize], )
        
        case func(Entity, *comp.Position, *comp.Velocity, ):
            fun(entity, &storePosition[entity.id >> entityPageBits][entity.id % entityPageSize], &storeVelocity[entity.id >> entityPageBits][entity.id % entityPageSize], )
        
        }
    }
}

func Select(selector interface{}) {
    match := true
    _ = match
    cont := true
    _ = cont
    var entity Entity
    for i := 0; i < currEntities; i++ {
        match = true
        entity = entities[i >> entityPageBits][i % entityPageSize]
        switch fun := selector.(type) {
        
        case func(Entity, *comp.BW, ) :
            match = match && entity.components[1] & 4096 != 0
            
            if match {
                fun(entity, &storeBW[entity.id >> entityPageBits][entity.id % entityPageSize], )
                
            }
        
        case func(Entity, *comp.Position, *comp.Velocity, ) :
            match = match && entity.components[0] & 1 != 0
            match = match && entity.components[0] & 2 != 0
            
            if match {
                fun(entity, &storePosition[entity.id >> entityPageBits][entity.id % entityPageSize], &storeVelocity[entity.id >> entityPageBits][entity.id % entityPageSize], )
                
            }
        
        case func(Entity):
            fun(entity)
        default:
            panic(fmt.Sprintf("unknown selector function: run go generate: %s", reflect.TypeOf(selector).String()))
        }
    }
}