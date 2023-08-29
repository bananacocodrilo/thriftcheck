 
namespace java com.bananacocodrilo.simple.service

enum ImplicitValues {
  ZERO,
  ONE,
  TWO,
}  

enum ImplicitValuesWithNegative {
  UNKNOWN =  -1,
  ZERO,
  ONE,
  TWO,
}  

enum ExplicitValues {
  ZERO = 0,
  ONE = 1,
  TWO = 2,
}  
 
enum ExplicitValuesWithNegative {
  UNKNOWN =  -1,
  ZERO = 0,
  ONE = 1,
  TWO = 2,
}  
 

struct StructWithoutZeroIndex {
    /**
     * This one is fine
     */
    1: optional string oneIndex
}

struct StructWithZeroIndex {
    /**
     * This should cause a problem as 0 index is not allowed
     */
    0: optional string zeroIndex
    /**
     * This one is fine
     */
     1: optional string oneIndex
} 



struct BADCaseStruct {
    /**
     * This one is fine
     */
     1: optional string oneIndex 
}

enum BADCaseEnum {
  ZERO = 0,
  ONE_OR_MORE = 1,
}

enum BadCaseEnumEntry {
  badenumentry = 0,
  Other_Bad_ENTRY = 1,
  BAD_ENTRY_3 = 1,
}

const string GOOD_CASE_CONST = "is good"
const string badCaseConst = "is bad"
 
service BADCaseService {
    void BADCaseMethod(1: bool BADCaseParam)
}

struct StructWithMap {
  1: optional map<string, string> normalMap
}

struct StructWithMapWithMap {
  1: optional map<string, map<string, string>> mapWithMap
}


