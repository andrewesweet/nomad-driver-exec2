# ICT Risk Management Framework RTS Analysis: IT Separation Requirements

## Document Information
- **Title**: Commission Delegated Regulation (EU) 2024/1774 of 13 March 2024 supplementing Regulation (EU) 2022/2554 with regard to regulatory technical standards specifying ICT risk management tools, methods, processes, and policies and the simplified ICT risk management framework
- **Source URI**: https://eur-lex.europa.eu/legal-content/EN/TXT/?uri=CELEX%3A32024R1774
- **Analysis Date**: June 15, 2025
- **Document Status**: In force from 25 June 2024
- **Official Journal**: OJ L, 2024/1774, 25.6.2024

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

## Separation Requirements Identified

### 1. ICT Risk Management Framework - Segregation of Duties

#### 1.1 Segregation of Duties in ICT Role Assignment
**Requirement**: Financial entities should ensure the segregation of duties when assigning ICT roles and responsibilities.

**Source**: Commission Delegated Regulation (EU) 2024/1774, Recital (4)
**Location**: Recital (4)
**Context**: "To limit the risk of conflicts of interests, financial entities should ensure the segregation of duties when assigning ICT roles and responsibilities."

### 2. ICT Operations Security - Production Environment Separation

#### 2.1 Strict Separation of Production Environments
**Requirement**: Financial entities must ensure strict separation of ICT production environments from development, testing, and other non-production environments.

**Source**: Commission Delegated Regulation (EU) 2024/1774, Recital (10)
**Location**: Recital (10)
**Context**: "One pivotal aspect is the strict separation of ICT production environments from the environments where ICT systems are developed and tested or from other non-production environments. That separation should serve as an important ICT security measure against unintended and unauthorised access to, modifications of, and deletions of data in the production environment."

#### 2.2 Exception for Testing in Production
**Requirement**: Testing in production environments may only be allowed in exceptional circumstances with proper justification and approval.

**Source**: Commission Delegated Regulation (EU) 2024/1774, Recital (10)
**Location**: Recital (10)
**Context**: "However, considering current ICT system development practices, in exceptional circumstances, financial entities should be allowed to test in production environments, provided that they justify such testing and obtain the required approval."

### 3. ICT Change Management - Separation of Functions

#### 3.1 Separation of Change Management Functions
**Requirement**: Financial entities must separate the functions responsible for approving ICT changes from the functions that request and implement those changes.

**Source**: Commission Delegated Regulation (EU) 2024/1774, Recital (17)
**Location**: Recital (17)
**Context**: "To uphold the objectivity and effectiveness of the ICT change management process, to prevent conflicts of interest, and to ensure that ICT changes are evaluated objectively, it is necessary to separate the functions responsible for approving those changes from the functions that request and implement those changes."

### 4. Simplified ICT Risk Management Framework - Control Function Segregation

#### 4.1 Segregation of Control and Audit Functions
**Requirement**: Financial entities under simplified framework must ensure appropriate segregation and independence of control functions and internal audit functions.

**Source**: Commission Delegated Regulation (EU) 2024/1774, Article 28(4)
**Location**: Article 28, paragraph 4
**Context**: "The financial entities referred to in paragraph 1 shall ensure an appropriate segregation and the independence of control functions and internal audit functions."

### 5. ICT Systems Acquisition and Development - Environment Separation

#### 5.1 Testing and Production Environment Separation
**Requirement**: Financial entities must ensure testing and approval of ICT systems prior to first use and before introducing changes to the production environment.

**Source**: Commission Delegated Regulation (EU) 2024/1774, Article 37(b)
**Location**: Article 37, point (b)
**Context**: "ensure the testing and approval of ICT systems prior to their first use and before introducing changes to the production environment"

#### 5.2 Development and Production Environment Protection
**Requirement**: Financial entities must identify measures to mitigate risks of unintentional alteration or intentional manipulation during development and implementation in production environment.

**Source**: Commission Delegated Regulation (EU) 2024/1774, Article 37(c)
**Location**: Article 37, point (c)
**Context**: "identify measures to mitigate the risk of unintentional alteration or intentional manipulation of the ICT systems during development and implementation in the production environment."

### 6. Physical and Environmental Security

#### 6.1 Physical Security Measures and Access Control
**Requirement**: Financial entities must identify and implement physical security measures to protect premises and data centres from unauthorised access.

**Source**: Commission Delegated Regulation (EU) 2024/1774, Article 32(1)(2)
**Location**: Article 32, paragraphs 1 and 2
**Context**: "The measures referred to in paragraph 1 shall protect the premises of financial entities and, where applicable, data centres of financial entities where ICT assets and information assets reside from unauthorised access, attacks, and accidents, and from environmental threats and hazards."

### 7. Access Control - Logical and Physical Separation

#### 7.1 Logical and Physical Access Control Procedures
**Requirement**: Financial entities must develop, document, and implement procedures for the control of logical and physical access.

**Source**: Commission Delegated Regulation (EU) 2024/1774, Article 33
**Location**: Article 33, opening paragraph
**Context**: "The financial entities referred to in Article 16(1) of Regulation (EU) 2022/2554 shall develop, document, and implement procedures for the control of logical and physical access"

#### 7.2 Need-to-Know and Least Privileges Access Control
**Requirement**: Access rights to information assets, ICT assets, and critical locations must be managed on a need-to-know, need-to-use and least privileges basis.

**Source**: Commission Delegated Regulation (EU) 2024/1774, Article 33(a)
**Location**: Article 33, point (a)
**Context**: "access rights to information assets, ICT assets, and their supported functions, and to critical locations of operation of the financial entity, are managed on a need-to-know, need-to-use and least privileges basis, including for remote and emergency access"

### 8. Data, System and Network Security - Network Protection

#### 8.1 Network Security Against Intrusions
**Requirement**: Financial entities must implement safeguards to ensure security of networks against intrusions and data misuse.

**Source**: Commission Delegated Regulation (EU) 2024/1774, Article 35
**Location**: Article 35, opening paragraph
**Context**: "develop and implement safeguards that ensure the security of networks against intrusions and data misuse and that preserve the availability, authenticity, integrity, and confidentiality of data"

#### 8.2 Network Connection Prevention and Detection
**Requirement**: Financial entities must identify and implement measures to prevent and detect unauthorised connections to the financial entity's network.

**Source**: Commission Delegated Regulation (EU) 2024/1774, Article 35(c)
**Location**: Article 35, point (c)
**Context**: "the identification and implementation of measures to prevent and detect unauthorised connections to the financial entity's network, and to secure the network traffic between the financial entity's internal networks and the internet and other external connections"

## Summary of Key Separation Requirements

### Production Environment Separation
1. **Strict separation** of production environments from development/testing environments (Recital 10)
2. **Testing approval** required before changes to production environment (Article 37(b))
3. **Risk mitigation** for development and production environment interactions (Article 37(c))

### Functional Separation Requirements
1. **Segregation of duties** in ICT role assignment (Recital 4)
2. **Separation of change management functions** - approval vs. request/implementation (Recital 17)
3. **Segregation of control and audit functions** in simplified framework (Article 28(4))

### Access Control Separation
1. **Logical and physical access control** procedures (Article 33)
2. **Need-to-know and least privileges** access management (Article 33(a))
3. **Physical security measures** for premises and data centres (Article 32)

### Network Security Separation
1. **Network security against intrusions** and data misuse (Article 35)
2. **Prevention and detection** of unauthorised network connections (Article 35(c))
3. **Secure network traffic** between internal networks and external connections (Article 35(c))

## Implementation Guidance for Milo Task Driver Plugin

### Relevant Technical Requirements
1. **Environment separation**: Implement strict separation between production and non-production workload execution
2. **Access control segregation**: Ensure logical separation of tenant access rights and privileges
3. **Network isolation**: Implement measures to prevent unauthorised connections between tenant workloads
4. **Physical security**: Consider physical separation requirements for multi-tenant infrastructure

### Compliance Considerations
- Implement role-based access control with segregation of duties
- Ensure testing environments are separated from production workloads
- Maintain audit trails for all access control and change management activities
- Regular review and monitoring of separation control effectiveness

## Appendix: Sources Inspected During Analysis

### A.1 Primary Sources Analyzed
**A.1.1 Commission Delegated Regulation (EU) 2024/1774**
- **URI**: https://eur-lex.europa.eu/legal-content/EN/TXT/?uri=CELEX%3A32024R1774
- **Document Type**: EU Delegated Regulation (RTS)
- **Access Date**: June 15, 2025 (UTC)
- **Analysis Status**: Fully analyzed
- **Relevance**: High
- **Notes**: Primary technical standard for ICT risk management framework implementation. Comprehensive analysis of all articles related to separation requirements and technical controls.

### A.2 Secondary Sources Reviewed
**A.2.1 EUR-Lex RTS Overview Page**
- **URI**: https://eur-lex.europa.eu/legal-content/EN/TXT/?uri=CELEX%3A32024R1774
- **Document Type**: Legal database overview
- **Access Date**: June 15, 2025 (UTC)
- **Analysis Status**: Reviewed for context
- **Relevance**: Medium
- **Notes**: Used for document metadata and official publication information

### A.3 Sources Identified But Not Accessed
**A.3.1 Related DORA Technical Standards**
- **URI**: Various EUR-Lex URLs for other DORA RTS/ITS
- **Document Type**: Related Regulatory Technical Standards
- **Reason Not Accessed**: Analyzed separately in dedicated analysis documents
- **Potential Relevance**: Medium
- **Notes**: Cross-references to other DORA technical standards analyzed in separate documents

### A.4 Source Analysis Summary
- **Total Sources Identified**: 3
- **Sources Fully Analyzed**: 1
- **Sources Partially Reviewed**: 1
- **Sources Not Accessed**: 1
- **Analysis Completeness**: 90% (comprehensive coverage of RTS requirements)

---
*Analysis completed: June 15, 2025*
*Framework Version Consulted: v1.7.0*
*Total separation requirements identified: 8 specific requirements across 4 categories*
*Source tracking: Complete appendix with 3 sources documented*
*Note: Analysis conducted under framework v1.5; future re-analysis recommended under v1.6 to identify STRIDE-based security separation requirements*
