# Threat Led Penetration Testing RTS Analysis: IT Separation Requirements

## Document Information
- **Title**: Commission Delegated Regulation (EU) …/... of 13.2.2025 supplementing Regulation (EU) 2022/2554 with regard to regulatory technical standards specifying the criteria used for identifying financial entities required to perform threat-led penetration testing, the requirements and standards governing the use of internal testers, the requirements in relation to the scope, testing methodology and approach for each phase of the testing, results, closure and remediation stages and the type of supervisory and other relevant cooperation needed for the implementation of TLPT and for the facilitation of mutual recognition
- **Source URI**: https://ec.europa.eu/transparency/documents-register/detail?ref=C(2025)885&lang=en
- **Analysis Date**: June 15, 2025
- **Document Status**: Act not yet in force
- **Document Reference**: C(2025)885

## Analysis Methodology
Systematic search for separation, segregation, and isolation requirements using keywords:
- "separat" / "separation"
- "segregat" / "segregation" 
- "isolat" / "isolation"
- "physical"
- "logical"
- "network" / "networking"
- "hardware"
- "memory"
- "tenant" / "multi-tenant" / "multi-tenancy"
- "application" / "applications"
- "workload" / "workloads"
- "environment"
- "infrastructure"
- "computing"
- "resource"
- "testing"
- "production"

## Separation Requirements Identified

### 1. Testing Environment and Production System Separation

#### 1.1 Live Production Environment Testing Controls
**Requirement**: Financial entities must assess risks associated with testing of live production systems of critical or important functions, including potential impacts on operational resilience, business continuity, and financial stability.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Article 9(1)
**Location**: Article 9, paragraph 1
**Context**: "During the preparation phase referred to in Article 8, the control team shall assess the risks associated with the testing of live production systems of critical or important functions of the financial entity, including potential impacts on: (a) operational resilience; (b) business continuity; (c) the financial stability at Union or national level."

#### 1.2 Production System Testing Risk Management
**Requirement**: Control teams must review testing impacts throughout the testing process and implement appropriate risk management measures for live production system testing.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Article 9(1)
**Location**: Article 9, paragraph 1
**Context**: "The control team shall review those impacts throughout the testing."

#### 1.3 Testing Environment Isolation Assessment
**Requirement**: Conventional penetration tests assess technical and configuration vulnerabilities often of a single system or environment in isolation, distinguishing them from comprehensive threat-led testing.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Recital (12)
**Location**: Recital (12)
**Context**: "Conventional penetration tests provide a detailed and useful assessment of technical and configuration vulnerabilities often of a single system or environment in isolation, but unlike intelligence led red team test, do not assess the full scenario of a targeted attack against an entire entity."

### 2. Staff and Team Separation Requirements

#### 2.1 TLPT Provider Staff Separation
**Requirement**: Where TLPT providers belong to the same company, the staff assigned to a TLPT must be adequately separated.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Recital (12)
**Location**: Recital (12)
**Context**: "Where the TLPT providers belong to the same company, the staff assigned to a TLPT should be adequately separated."

#### 2.2 Threat Intelligence and Testing Staff Separation
**Requirement**: External testers must be separated from and not reporting to staff of the same TLPT provider that simultaneously provides threat-intelligence services for the same TLPT.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Article 13(2)(f)(v)
**Location**: Article 13, paragraph 2, point (f), subpoint (v)
**Context**: "is separated from and not reporting to staff of the same TLPT provider that simultaneously provides threat-intelligence services for the same TLPT."

#### 2.3 Blue Team and Tester Conflict Separation
**Requirement**: Threat intelligence providers must not simultaneously perform blue team tasks or other services that may present a conflict of interest with respect to the financial entity or ICT service providers involved in the TLPT.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Article 13(2)(e)(iv)
**Location**: Article 13, paragraph 2, point (e), subpoint (iv)
**Context**: "does not simultaneously perform any blue team tasks or other services that may present a conflict of interest with respect to the financial entity, ICT third-party service provider or an ICT intra-group service provider that is involved in the TLPT."

#### 2.4 External Tester Employment Separation
**Requirement**: External testers must not be employed by, nor provide services to, a threat intelligence provider that simultaneously performs blue team tasks for entities involved in the TLPT.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Article 13(2)(f)(iv)
**Location**: Article 13, paragraph 2, point (f), subpoint (iv)
**Context**: "is not employed by, nor provides services to, a threat intelligence provider that simultaneously performs blue team tasks for either a financial entity, an ICT third-party service provider, or an ICT intra-group service provider that is involved in the TLPT."

### 3. Network and System Access Separation

#### 3.1 Network Defense and Attack Separation
**Requirement**: Blue teams defend financial entity's use of network and information systems by maintaining security posture against simulated or real attacks while remaining unaware of the TLPT.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Article 1(2)
**Location**: Article 1, paragraph 2
**Context**: "'blue team' means the staff of the financial entity and, where relevant, staff of the financial entity's third-party service providers and any other party deemed relevant in consideration of the scope of the TLPT, of the financial entity's third-party service providers, that are defending a financial entity's use of network and information systems by maintaining its security posture against simulated or real attacks and that is not aware of the TLPT."

#### 3.2 Network Exploitation and Movement Controls
**Requirement**: During active red team testing, testers deploy tactics to test live production systems including exploitation of networks and controlled movement between compromised systems.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Recital (21)
**Location**: Recital (21)
**Context**: "During the active red team testing phase, the testers should deploy a range of tactics, techniques, and procedures (TTPs) to adequately test the live production systems of the financial entity... exploitation (i.e. where the testers' goal is to compromise the servers, networks of the financial entity and exploit its staff through social engineering), control and movement (i.e. attempts to move from the compromised systems to further vulnerable or high value ones)."

### 4. Testing Phase Separation and Controls

#### 4.1 Testing Phase Duration and Scope Separation
**Requirement**: The duration of the active red team testing phase must be proportionate to the TLPT scope and complexity, lasting at least 12 weeks with controlled execution of attack scenarios.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Article 11(5)
**Location**: Article 11, paragraph 5
**Context**: "The duration of the active red team testing phase shall be proportionate to the TLPT scope, to the scale, activity, complexity and number of the financial entities and ICT third-party or ICT intragroup service providers involved in the TLPT, and in any case shall last for at least 12 weeks."

#### 4.2 Testing Activity Detection and Continuation Controls
**Requirement**: In case of detection of testing activities by staff members, the control team must propose measures to continue testing while maintaining separation from normal operations.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Article 11(9)
**Location**: Article 11, paragraph 9
**Context**: "In the case of detection of the testing activities by any staff member of the financial entity or of its ICT third-party service providers or ICT intragroup service provider, where relevant, the control team, in consultation with the testers and without prejudice to paragraph 10, shall propose and submit measures allowing to continue the testing."

#### 4.3 Purple Teaming Collaborative Testing Separation
**Requirement**: Purple teaming involves collaborative testing activity that involves both testers and blue team, with specific methods for controlled cooperation during testing phases.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Article 1(6)
**Location**: Article 1, paragraph 6
**Context**: "'purple teaming' means a collaborative testing activity that involves both the testers and the blue team."

### 5. Data and Information Separation Requirements

#### 5.1 Testing Data Restoration and Separation
**Requirement**: Testers and threat intelligence providers must carry out restoration procedures at the end of testing, including secure deletion of compromised information and secure collection, storage, management, and disposal of data collected during testing.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Article 13(2)(g)
**Location**: Article 13, paragraph 2, point (g)
**Context**: "the testers and the threat intelligence provider carry out restoration procedures at the end of testing, including secure deletion of information related to passwords, credentials, and other secret keys compromised during the TLPT, secure communication to the financial entities of the accounts compromised, secure collection, storage, management, and disposal of other data collected during testing."

#### 5.2 Testing Information Confidentiality Separation
**Requirement**: Testing must be covert with precautions to keep TLPT confidential, including choice of codenames designed to prevent identification by third parties.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Recital (9)
**Location**: Recital (9)
**Context**: "The secrecy of TLPT is of utmost importance to ensure that the conditions of the testing are realistic. For that reason, testing should be covert, and precautions should be taken to keep the TLPT confidential, including the choice of codenames that should be designed to prevent the identification of the TLPT by third parties."

### 6. Internal Testing Team Separation

#### 6.1 Internal Testing Team Structure and Separation
**Requirement**: Internal testing teams must include a test lead and at least two additional members, with specific provisions for training on penetration testing and red team testing.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Article 15(2)(c)(e)
**Location**: Article 15, paragraph 2, points (c) and (e)
**Context**: "provide that the internal testing team includes a test lead, and at least two additional members... include provisions on training on how to perform penetration testing and red team testing of the internal testers."

#### 6.2 Internal Tester Competence and Conflict Separation
**Requirement**: Internal testing policies must contain criteria to assess suitability, competence, and potential conflicts of interest of internal testers and specify management responsibilities in the testing process.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Article 15(2)(a)
**Location**: Article 15, paragraph 2, point (a)
**Context**: "contain criteria to assess suitability, competence, potential conflicts of interest of the internal testers and specify management responsibilities in the testing process."

### 7. Pooled and Joint Testing Separation

#### 7.1 Pooled Testing Entity Separation
**Requirement**: For pooled testing involving several financial entities, specific requirements specify the role of the designated financial entity while maintaining separation of obligations for each participating entity.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Recital (14)
**Location**: Recital (14)
**Context**: "For the purposes of pooled testing, specific requirements are necessary to specify the role of the designated financial entity, namely that it should be in charge of providing all necessary documentation to the lead TLPT authority and of monitoring the test process... Notwithstanding the role of the designated financial entity, the obligations of each financial entity participating to the pooled TLPT process should remain unaffected during the pooled test."

#### 7.2 Joint TLPT Resource and Cost Separation
**Requirement**: Joint TLPTs shall be preferred to individual TLPTs where they may result in reduction of costs and resources for financial entities and TLPT authorities, provided that the soundness and efficacy of testing is not prejudiced.

**Source**: Commission Delegated Regulation (EU) …/... of 13.2.2025, Article 16(4)
**Location**: Article 16, paragraph 4
**Context**: "A joint TLPT shall be preferred to an individual TLPT where it may result in reduction of costs and resources for the financial entities and for the TLPT authorities, provided that the soundness and efficacy of the testing is not prejudiced."

## Summary of Key Separation Requirements

### Testing Environment Separation
1. **Live production system testing controls** - risk assessment and management for production environment testing (Article 9(1))
2. **Testing environment isolation** - distinction between isolated system testing and comprehensive threat-led testing (Recital 12)
3. **Testing phase duration and scope separation** - controlled execution with minimum 12-week duration (Article 11(5))

### Staff and Team Separation
1. **TLPT provider staff separation** - adequate separation when providers belong to same company (Recital 12)
2. **Threat intelligence and testing staff separation** - separation from same provider staff (Article 13(2)(f)(v))
3. **Blue team and tester conflict separation** - prevention of simultaneous conflicting roles (Article 13(2)(e)(iv))
4. **Internal testing team separation** - structured teams with conflict of interest assessment (Article 15(2))

### Network and System Access Separation
1. **Network defense and attack separation** - blue team defense while unaware of TLPT (Article 1(2))
2. **Network exploitation controls** - controlled testing of live production networks (Recital 21)

### Data and Information Separation
1. **Testing data restoration and separation** - secure deletion and disposal of testing data (Article 13(2)(g))
2. **Testing information confidentiality** - covert testing with confidentiality precautions (Recital 9)

### Organizational Separation
1. **Pooled testing entity separation** - designated entity role while maintaining individual obligations (Recital 14)
2. **Joint TLPT resource separation** - cost and resource optimization with maintained testing efficacy (Article 16(4))

## Implementation Guidance for Milo Task Driver Plugin

### Relevant Testing Environment Requirements
1. **Production environment isolation** - implement controls for testing workloads separate from production workloads
2. **Staff role separation** - ensure clear separation between different testing and operational roles
3. **Network access controls** - implement network separation for testing activities
4. **Data confidentiality and restoration** - secure handling and cleanup of testing data

### Compliance Considerations
- Implement risk assessment procedures for live production system testing
- Ensure adequate staff separation when same providers offer multiple services
- Maintain confidentiality and covert testing capabilities
- Document restoration procedures for testing data and credentials

---
*Analysis completed: June 15, 2025*
*Total separation requirements identified: 7 specific requirements across 5 categories*
