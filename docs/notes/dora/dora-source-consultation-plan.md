# DORA Source Consultation Plan

## 1. Planning Information
- **Plan Version**: 1.0.0
- **Created**: 2025-06-15 22:40:15 UTC
- **Framework Version**: Meta-Regulatory Analysis Framework v1.8.0
- **Analysis Target**: Digital Operational Resilience Act (DORA) - EU Regulation 2022/2554 and Technical Standards
- **Requested By**: Andrew Sweet (andrew.sweet@cantab.net)
- **Plan Status**: APPROVED
- **Approved By**: Andrew Sweet (andrew.sweet@cantab.net)
- **Approval Date**: 2025-06-15 22:43:20 UTC
- **Approval Method**: Chat message confirmation

## 2. Analysis Scope Definition

### 2.1 Primary Analysis Objectives
- Identify all IT separation, segregation, and isolation requirements in DORA regulation and technical standards
- Analyze threat actor specifications and security separation requirements using STRIDE threat model
- Assess multi-tenant relevance for containerized financial services infrastructure
- Extract technical implementation requirements for Milo Nomad Task Driver Plugin
- Create comprehensive requirements matrix for regulatory compliance

### 2.2 Scope Exclusions
- Human-to-machine interactions (per framework v1.8.0)
- Segregation of duties for human users
- Privileged access management for humans
- Organizational governance processes (focus on technical infrastructure)
- Non-technical compliance procedures

## 3. Source Identification Strategy

### 3.1 Primary Sources (Tier 1)
1. **Regulation (EU) 2022/2554** - DORA Main Regulation
   - URI: https://eur-lex.europa.eu/eli/reg/2022/2554/oj
   - Status: In force from 17 January 2025
   - Priority: Critical - Foundation document

2. **Commission Delegated Regulation (EU) 2024/1774** - ICT Risk Management Framework RTS
   - URI: https://eur-lex.europa.eu/legal-content/EN/TXT/?uri=CELEX%3A32024R1774
   - Status: Published, applicable from 17 January 2025
   - Priority: High - Technical implementation details

3. **Commission Delegated Regulation (EU) 2025/885** - Threat-Led Penetration Testing RTS
   - URI: https://ec.europa.eu/transparency/documents-register/detail?ref=C(2025)885&lang=en
   - Status: Recently adopted
   - Priority: High - Testing environment separation requirements

4. **Commission Delegated Regulation (EU) C(2025)1682** - Subcontracting RTS
   - URI: https://ec.europa.eu/transparency/documents-register/detail?ref=C(2025)1682&lang=en
   - Status: Recently adopted
   - Priority: High - Third-party segregation requirements

### 3.2 Secondary Sources (Tier 2)
1. **ICT Third-Party Risk Management RTS**
   - URI: [To be identified from official EU sources]
   - Status: Under development/recently published
   - Priority: High - Multi-tenancy and cloud segregation

2. **EBA Guidelines on ICT and Security Risk Management (EBA/GL/2019/04)**
   - URI: https://www.eba.europa.eu/regulation-and-policy/internal-governance/guidelines-on-ict-and-security-risk-management
   - Status: Superseded by DORA but relevant for context
   - Priority: Medium - Historical context and interpretation

3. **EBA Guidelines on Outsourcing Arrangements (EBA/GL/2019/02)**
   - URI: https://www.eba.europa.eu/regulation-and-policy/internal-governance/guidelines-on-outsourcing-arrangements
   - Status: Active, complementary to DORA
   - Priority: Medium - Cloud and outsourcing segregation

### 3.3 Tertiary Sources (Tier 3)
1. **EIOPA DORA Information Page**
   - URI: https://www.eiopa.europa.eu/digital-operational-resilience-act-dora_en
   - Priority: Low - Supervisory interpretation

2. **ESMA DORA Implementation Guidance**
   - URI: https://www.esma.europa.eu/ (specific DORA pages)
   - Priority: Low - Market infrastructure perspective

3. **ECB Supervisory Expectations for IT Risk**
   - URI: https://www.bankingsupervision.europa.eu/
   - Priority: Low - Banking supervision context

## 4. Source Access and Acquisition Plan

### 4.1 Publicly Available Sources
- **EUR-Lex Database**: Primary access point for all EU regulations and delegated acts
- **European Supervisory Authorities**: EBA, EIOPA, ESMA official websites
- **European Commission**: Transparency register for recent delegated acts
- **ECB Banking Supervision**: Supervisory guidance and expectations

### 4.2 Subscription-Required Sources
- None identified - all primary sources are publicly available through official EU channels

### 4.3 Restricted Access Sources
- None identified - regulatory transparency requirements ensure public access to all relevant documents

## 5. Analysis Methodology

### 5.1 Keyword Search Strategy

**Traditional Separation Keywords**:
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

**STRIDE-Enhanced Keywords**:
- "spoofing" / "authentication" / "identity"
- "tampering" / "integrity" / "modification"
- "repudiation" / "audit" / "logging"
- "disclosure" / "confidentiality" / "privacy"
- "denial" / "availability" / "service"
- "elevation" / "privilege" / "authorization"

**DORA-Specific Keywords**:
- "ICT risk management"
- "operational resilience"
- "incident management"
- "digital operational resilience testing"
- "third-party risk"
- "critical ICT third-party service provider"
- "concentration risk"
- "business continuity"
- "recovery"

### 5.2 Source Attribution Requirements
- **UTC timestamp format**: YYYY-MM-DD HH:MM:SS UTC
- **Complete URI documentation**: Permalinks using EUR-Lex identifiers where available
- **Document version tracking**: Official publication dates and version numbers
- **Access method recording**: Direct download, web access, database query
- **Retrieval verification**: Document hash or size verification where possible

## 6. Expected Deliverables

### 6.1 Analysis Documents
1. **dora-comprehensive-analysis-v3.md** - Complete DORA analysis replacing existing v2.0
2. **dora-requirements-matrix.md** - Consolidated requirements across all DORA sources
3. **dora-implementation-guidance.md** - Technical implementation recommendations for Milo plugin

### 6.2 Requirements Matrix
- **Separation Categories**: Physical, Logical, Environment, Multi-Tenant, Network
- **STRIDE Mapping**: Requirements mapped to STRIDE threat categories
- **Priority Classification**: Critical, High, Medium, Low based on regulatory language
- **Implementation Complexity**: Technical difficulty assessment for each requirement

### 6.3 Implementation Guidance
- **Container Orchestration**: Kubernetes/Nomad-specific separation controls
- **Network Policies**: Traffic segregation and isolation requirements
- **Storage Segregation**: Data isolation and tenant separation
- **Resource Management**: CPU, memory, and I/O isolation requirements
- **Security Controls**: Authentication, authorization, and audit separation

## 7. Resource Requirements and Timeline

### 7.1 Estimated Analysis Duration
- **Phase 1 - Source Acquisition**: 2-3 hours (document download and verification)
- **Phase 2 - Primary Analysis**: 6-8 hours (main regulation and key RTS)
- **Phase 3 - Technical Standards**: 4-6 hours (remaining RTS and implementation details)
- **Phase 4 - Consolidation**: 3-4 hours (requirements matrix and implementation guidance)
- **Total Estimated Duration**: 15-21 hours

### 7.2 Source Access Requirements
- **Internet Access**: Required for EUR-Lex and regulatory authority websites
- **Browser Capabilities**: PDF viewing and download capabilities
- **Storage Space**: Approximately 50-100MB for document storage
- **No Special Credentials**: All sources publicly accessible

### 7.3 Dependencies
- **Framework Compliance**: Must follow Meta-Regulatory Analysis Framework v1.8.0
- **User Approval**: This plan must be approved before analysis commences
- **Source Availability**: Dependent on EUR-Lex and regulatory websites being accessible
- **Document Stability**: Recent delegated acts may have updates or corrections

## 8. Risk Assessment

### 8.1 Source Availability Risks
- **Low Risk**: EUR-Lex provides stable, permanent access to EU legislation
- **Mitigation**: Multiple access points (direct EUR-Lex, regulatory authority websites)
- **Contingency**: Archive.org or cached versions if primary sources temporarily unavailable

### 8.2 Scope Creep Risks
- **Medium Risk**: DORA technical standards are extensive and interconnected
- **Mitigation**: Strict adherence to IT separation focus, exclude purely procedural requirements
- **Control Mechanism**: Regular scope validation against approved plan objectives

### 8.3 Quality Assurance Measures
- **Source Verification**: Cross-reference multiple official sources for consistency
- **Keyword Validation**: Systematic application of approved keyword methodology
- **Requirements Traceability**: Each requirement linked to specific regulatory citation
- **STRIDE Mapping Validation**: Technical review of threat model application

## Appendix: Preliminary Source Inventory

### A.1 Confirmed Available Sources
1. **Regulation (EU) 2022/2554** - 79 pages, PDF available
2. **Commission Delegated Regulation (EU) 2024/1774** - Confirmed published
3. **EBA Guidelines EBA/GL/2019/04** - 87 pages, superseded but contextually relevant

### A.2 Sources Requiring Verification
1. **Commission Delegated Regulation (EU) 2025/885** - Recently adopted, availability to be confirmed
2. **Commission Delegated Regulation (EU) C(2025)1682** - Recently adopted, availability to be confirmed
3. **ICT Third-Party Risk Management RTS** - Publication status to be verified

### A.3 Source Access Methods
- **Primary Method**: Direct download from EUR-Lex using CELEX numbers
- **Secondary Method**: Regulatory authority websites (EBA, EIOPA, ESMA)
- **Tertiary Method**: European Commission transparency register

---
*Plan created: 2025-06-15 22:40:15 UTC*
*Framework version: v1.8.0*
*Status: PENDING_APPROVAL - User review and confirmation required before proceeding*
