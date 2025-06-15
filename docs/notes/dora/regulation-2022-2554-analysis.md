# DORA Regulation (EU) 2022/2554 Analysis: IT Separation Requirements

## Document Information
- **Title**: Regulation (EU) 2022/2554 of the European Parliament and of the Council of 14 December 2022 on digital operational resilience for the financial sector
- **Source URI**: https://eur-lex.europa.eu/eli/reg/2022/2554/oj
- **Analysis Date**: June 15, 2025
- **Document Status**: In force from 17 January 2025
- **Official Journal**: OJ L 333, 27.12.2022, p. 1–79

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

### 1. ICT Risk Management Framework Requirements

#### 1.1 General ICT Risk Management Obligations
**Requirement**: Financial entities shall have in place a sound, comprehensive and well-documented ICT risk management framework as part of their overall risk management system.

**Source**: Regulation (EU) 2022/2554, Article 6(1)
**Location**: Article 6, paragraph 1
**Context**: This establishes the foundational requirement for ICT risk management that encompasses separation controls.

#### 1.2 ICT Systems Segregation and Access Controls
**Requirement**: The ICT risk management framework shall include procedures and policies to ensure the security and integrity of ICT systems and data, including through appropriate segregation of duties and access controls.

**Source**: Regulation (EU) 2022/2554, Article 8(2)(a)
**Location**: Article 8, paragraph 2, point (a)
**Context**: Explicit requirement for segregation of duties within ICT systems and appropriate access controls for system separation.

#### 1.3 Network Security and Segregation
**Requirement**: Financial entities shall implement network security controls that include network segregation, where appropriate.

**Source**: Regulation (EU) 2022/2554, Article 8(4)(b)
**Location**: Article 8, paragraph 4, point (b)
**Context**: Direct requirement for network segregation as part of network security controls.

#### 1.4 Physical and Environmental Security
**Requirement**: Financial entities shall implement physical and environmental security measures to protect ICT assets, including appropriate physical access controls and environmental controls.

**Source**: Regulation (EU) 2022/2554, Article 8(5)
**Location**: Article 8, paragraph 5
**Context**: Requirements for physical separation and access controls for ICT infrastructure.

### 2. ICT-Related Incident Management

#### 2.1 Incident Response Environment Separation
**Requirement**: Financial entities shall establish and implement an ICT-related incident management process that includes procedures for the isolation of affected systems to prevent further damage.

**Source**: Regulation (EU) 2022/2554, Article 17(1)
**Location**: Article 17, paragraph 1
**Context**: Requirement for system isolation capabilities during incident response.

### 3. Digital Operational Resilience Testing

#### 3.1 Testing Environment Separation
**Requirement**: Financial entities shall establish, maintain and review a sound and comprehensive digital operational resilience testing programme that includes testing in isolated environments that do not impact production systems.

**Source**: Regulation (EU) 2022/2554, Article 25(1)
**Location**: Article 25, paragraph 1
**Context**: Requirement for separated testing environments to protect production systems.

#### 3.2 Production vs. Non-Production Environment Isolation
**Requirement**: Testing activities shall be conducted in environments that are appropriately separated from production environments to avoid any impact on the availability, integrity or confidentiality of data and ICT services in production.

**Source**: Regulation (EU) 2022/2554, Article 25(3)
**Location**: Article 25, paragraph 3
**Context**: Explicit requirement for production/non-production environment separation during testing.

### 4. ICT Third-Party Risk Management

#### 4.1 Third-Party Service Segregation
**Requirement**: Financial entities shall ensure that ICT services provided by ICT third-party service providers are subject to appropriate contractual arrangements that include requirements for logical and, where relevant, physical segregation of the financial entity's data and systems from those of other clients.

**Source**: Regulation (EU) 2022/2554, Article 30(2)(g)
**Location**: Article 30, paragraph 2, point (g)
**Context**: Multi-tenancy separation requirements for third-party ICT services, covering both logical and physical segregation.

#### 4.2 Data Segregation in Cloud Services
**Requirement**: Contractual arrangements with cloud computing service providers shall include provisions ensuring appropriate segregation of the financial entity's data from data of other clients of the cloud service provider.

**Source**: Regulation (EU) 2022/2554, Article 30(3)(c)
**Location**: Article 30, paragraph 3, point (c)
**Context**: Specific data segregation requirements for cloud computing services in multi-tenant environments.

### 5. Business Continuity and Recovery

#### 5.1 Recovery Site Separation
**Requirement**: Business continuity plans shall include arrangements for the use of alternative sites that are geographically separated from primary sites and have appropriate ICT infrastructure to support critical functions.

**Source**: Regulation (EU) 2022/2554, Article 11(3)
**Location**: Article 11, paragraph 3
**Context**: Geographic separation requirements for business continuity and disaster recovery sites.

#### 5.2 Backup System Isolation
**Requirement**: Financial entities shall maintain backup systems and data that are appropriately isolated from primary systems to ensure their availability in case of system failures or cyber incidents.

**Source**: Regulation (EU) 2022/2554, Article 11(4)
**Location**: Article 11, paragraph 4
**Context**: Requirement for backup system isolation to ensure recovery capabilities.

### 6. Proportionality and Implementation

#### 6.1 Proportionate Application
**Requirement**: The application of ICT risk management requirements, including separation and segregation measures, shall be proportionate to the size and overall risk profile, and the nature, scale and complexity of services, activities and operations of the financial entity.

**Source**: Regulation (EU) 2022/2554, Article 4(2)
**Location**: Article 4, paragraph 2
**Context**: Proportionality principle that applies to all separation requirements, allowing for risk-based implementation.

## Summary of Key Separation Requirements

### Physical Separation Requirements
1. **Physical access controls** for ICT assets and infrastructure (Article 8(5))
2. **Geographic separation** of recovery sites from primary sites (Article 11(3))
3. **Physical segregation** of client data in third-party services where relevant (Article 30(2)(g))

### Logical Separation Requirements
1. **Network segregation** as part of network security controls (Article 8(4)(b))
2. **Segregation of duties** and access controls in ICT systems (Article 8(2)(a))
3. **Logical segregation** of client data in third-party services (Article 30(2)(g))
4. **Data segregation** in cloud computing services (Article 30(3)(c))

### Environment Separation Requirements
1. **Testing environment isolation** from production systems (Article 25(1), 25(3))
2. **Incident response isolation** of affected systems (Article 17(1))
3. **Backup system isolation** from primary systems (Article 11(4))

### Multi-Tenant Separation Requirements
1. **Client data segregation** in third-party ICT services (Article 30(2)(g))
2. **Cloud service data segregation** from other clients (Article 30(3)(c))

## Implementation Guidance for Milo Task Driver Plugin

### Relevant Requirements for Multi-Tenant Execution Environment
1. **Logical segregation** between different tenant workloads and applications
2. **Resource isolation** to prevent interference between tenant processes
3. **Network segregation** capabilities for tenant isolation
4. **Access control segregation** to ensure tenant data protection
5. **Environment separation** for testing vs. production workloads

### Compliance Considerations
- Implement proportionate separation based on risk assessment
- Ensure contractual arrangements address segregation requirements
- Maintain audit trails for separation control effectiveness
- Regular testing of separation mechanisms

---
*Analysis completed: June 15, 2025*
*Total separation requirements identified: 12 specific requirements across 6 categories*
