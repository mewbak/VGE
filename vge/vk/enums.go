package vk

type Format int32

const (
	FORMATUndefined                            = Format(0)
	FORMATR4g4UnormPack8                       = Format(1)
	FORMATR4g4b4a4UnormPack16                  = Format(2)
	FORMATB4g4r4a4UnormPack16                  = Format(3)
	FORMATR5g6b5UnormPack16                    = Format(4)
	FORMATB5g6r5UnormPack16                    = Format(5)
	FORMATR5g5b5a1UnormPack16                  = Format(6)
	FORMATB5g5r5a1UnormPack16                  = Format(7)
	FORMATA1r5g5b5UnormPack16                  = Format(8)
	FORMATR8Unorm                              = Format(9)
	FORMATR8Snorm                              = Format(10)
	FORMATR8Uscaled                            = Format(11)
	FORMATR8Sscaled                            = Format(12)
	FORMATR8Uint                               = Format(13)
	FORMATR8Sint                               = Format(14)
	FORMATR8Srgb                               = Format(15)
	FORMATR8g8Unorm                            = Format(16)
	FORMATR8g8Snorm                            = Format(17)
	FORMATR8g8Uscaled                          = Format(18)
	FORMATR8g8Sscaled                          = Format(19)
	FORMATR8g8Uint                             = Format(20)
	FORMATR8g8Sint                             = Format(21)
	FORMATR8g8Srgb                             = Format(22)
	FORMATR8g8b8Unorm                          = Format(23)
	FORMATR8g8b8Snorm                          = Format(24)
	FORMATR8g8b8Uscaled                        = Format(25)
	FORMATR8g8b8Sscaled                        = Format(26)
	FORMATR8g8b8Uint                           = Format(27)
	FORMATR8g8b8Sint                           = Format(28)
	FORMATR8g8b8Srgb                           = Format(29)
	FORMATB8g8r8Unorm                          = Format(30)
	FORMATB8g8r8Snorm                          = Format(31)
	FORMATB8g8r8Uscaled                        = Format(32)
	FORMATB8g8r8Sscaled                        = Format(33)
	FORMATB8g8r8Uint                           = Format(34)
	FORMATB8g8r8Sint                           = Format(35)
	FORMATB8g8r8Srgb                           = Format(36)
	FORMATR8g8b8a8Unorm                        = Format(37)
	FORMATR8g8b8a8Snorm                        = Format(38)
	FORMATR8g8b8a8Uscaled                      = Format(39)
	FORMATR8g8b8a8Sscaled                      = Format(40)
	FORMATR8g8b8a8Uint                         = Format(41)
	FORMATR8g8b8a8Sint                         = Format(42)
	FORMATR8g8b8a8Srgb                         = Format(43)
	FORMATB8g8r8a8Unorm                        = Format(44)
	FORMATB8g8r8a8Snorm                        = Format(45)
	FORMATB8g8r8a8Uscaled                      = Format(46)
	FORMATB8g8r8a8Sscaled                      = Format(47)
	FORMATB8g8r8a8Uint                         = Format(48)
	FORMATB8g8r8a8Sint                         = Format(49)
	FORMATB8g8r8a8Srgb                         = Format(50)
	FORMATA8b8g8r8UnormPack32                  = Format(51)
	FORMATA8b8g8r8SnormPack32                  = Format(52)
	FORMATA8b8g8r8UscaledPack32                = Format(53)
	FORMATA8b8g8r8SscaledPack32                = Format(54)
	FORMATA8b8g8r8UintPack32                   = Format(55)
	FORMATA8b8g8r8SintPack32                   = Format(56)
	FORMATA8b8g8r8SrgbPack32                   = Format(57)
	FORMATA2r10g10b10UnormPack32               = Format(58)
	FORMATA2r10g10b10SnormPack32               = Format(59)
	FORMATA2r10g10b10UscaledPack32             = Format(60)
	FORMATA2r10g10b10SscaledPack32             = Format(61)
	FORMATA2r10g10b10UintPack32                = Format(62)
	FORMATA2r10g10b10SintPack32                = Format(63)
	FORMATA2b10g10r10UnormPack32               = Format(64)
	FORMATA2b10g10r10SnormPack32               = Format(65)
	FORMATA2b10g10r10UscaledPack32             = Format(66)
	FORMATA2b10g10r10SscaledPack32             = Format(67)
	FORMATA2b10g10r10UintPack32                = Format(68)
	FORMATA2b10g10r10SintPack32                = Format(69)
	FORMATR16Unorm                             = Format(70)
	FORMATR16Snorm                             = Format(71)
	FORMATR16Uscaled                           = Format(72)
	FORMATR16Sscaled                           = Format(73)
	FORMATR16Uint                              = Format(74)
	FORMATR16Sint                              = Format(75)
	FORMATR16Sfloat                            = Format(76)
	FORMATR16g16Unorm                          = Format(77)
	FORMATR16g16Snorm                          = Format(78)
	FORMATR16g16Uscaled                        = Format(79)
	FORMATR16g16Sscaled                        = Format(80)
	FORMATR16g16Uint                           = Format(81)
	FORMATR16g16Sint                           = Format(82)
	FORMATR16g16Sfloat                         = Format(83)
	FORMATR16g16b16Unorm                       = Format(84)
	FORMATR16g16b16Snorm                       = Format(85)
	FORMATR16g16b16Uscaled                     = Format(86)
	FORMATR16g16b16Sscaled                     = Format(87)
	FORMATR16g16b16Uint                        = Format(88)
	FORMATR16g16b16Sint                        = Format(89)
	FORMATR16g16b16Sfloat                      = Format(90)
	FORMATR16g16b16a16Unorm                    = Format(91)
	FORMATR16g16b16a16Snorm                    = Format(92)
	FORMATR16g16b16a16Uscaled                  = Format(93)
	FORMATR16g16b16a16Sscaled                  = Format(94)
	FORMATR16g16b16a16Uint                     = Format(95)
	FORMATR16g16b16a16Sint                     = Format(96)
	FORMATR16g16b16a16Sfloat                   = Format(97)
	FORMATR32Uint                              = Format(98)
	FORMATR32Sint                              = Format(99)
	FORMATR32Sfloat                            = Format(100)
	FORMATR32g32Uint                           = Format(101)
	FORMATR32g32Sint                           = Format(102)
	FORMATR32g32Sfloat                         = Format(103)
	FORMATR32g32b32Uint                        = Format(104)
	FORMATR32g32b32Sint                        = Format(105)
	FORMATR32g32b32Sfloat                      = Format(106)
	FORMATR32g32b32a32Uint                     = Format(107)
	FORMATR32g32b32a32Sint                     = Format(108)
	FORMATR32g32b32a32Sfloat                   = Format(109)
	FORMATR64Uint                              = Format(110)
	FORMATR64Sint                              = Format(111)
	FORMATR64Sfloat                            = Format(112)
	FORMATR64g64Uint                           = Format(113)
	FORMATR64g64Sint                           = Format(114)
	FORMATR64g64Sfloat                         = Format(115)
	FORMATR64g64b64Uint                        = Format(116)
	FORMATR64g64b64Sint                        = Format(117)
	FORMATR64g64b64Sfloat                      = Format(118)
	FORMATR64g64b64a64Uint                     = Format(119)
	FORMATR64g64b64a64Sint                     = Format(120)
	FORMATR64g64b64a64Sfloat                   = Format(121)
	FORMATB10g11r11UfloatPack32                = Format(122)
	FORMATE5b9g9r9UfloatPack32                 = Format(123)
	FORMATD16Unorm                             = Format(124)
	FORMATX8D24UnormPack32                     = Format(125)
	FORMATD32Sfloat                            = Format(126)
	FORMATS8Uint                               = Format(127)
	FORMATD16UnormS8Uint                       = Format(128)
	FORMATD24UnormS8Uint                       = Format(129)
	FORMATD32SfloatS8Uint                      = Format(130)
	FORMATBc1RgbUnormBlock                     = Format(131)
	FORMATBc1RgbSrgbBlock                      = Format(132)
	FORMATBc1RgbaUnormBlock                    = Format(133)
	FORMATBc1RgbaSrgbBlock                     = Format(134)
	FORMATBc2UnormBlock                        = Format(135)
	FORMATBc2SrgbBlock                         = Format(136)
	FORMATBc3UnormBlock                        = Format(137)
	FORMATBc3SrgbBlock                         = Format(138)
	FORMATBc4UnormBlock                        = Format(139)
	FORMATBc4SnormBlock                        = Format(140)
	FORMATBc5UnormBlock                        = Format(141)
	FORMATBc5SnormBlock                        = Format(142)
	FORMATBc6hUfloatBlock                      = Format(143)
	FORMATBc6hSfloatBlock                      = Format(144)
	FORMATBc7UnormBlock                        = Format(145)
	FORMATBc7SrgbBlock                         = Format(146)
	FORMATEtc2R8g8b8UnormBlock                 = Format(147)
	FORMATEtc2R8g8b8SrgbBlock                  = Format(148)
	FORMATEtc2R8g8b8a1UnormBlock               = Format(149)
	FORMATEtc2R8g8b8a1SrgbBlock                = Format(150)
	FORMATEtc2R8g8b8a8UnormBlock               = Format(151)
	FORMATEtc2R8g8b8a8SrgbBlock                = Format(152)
	FORMATEacR11UnormBlock                     = Format(153)
	FORMATEacR11SnormBlock                     = Format(154)
	FORMATEacR11g11UnormBlock                  = Format(155)
	FORMATEacR11g11SnormBlock                  = Format(156)
	FORMATAstc4x4UnormBlock                    = Format(157)
	FORMATAstc4x4SrgbBlock                     = Format(158)
	FORMATAstc5x4UnormBlock                    = Format(159)
	FORMATAstc5x4SrgbBlock                     = Format(160)
	FORMATAstc5x5UnormBlock                    = Format(161)
	FORMATAstc5x5SrgbBlock                     = Format(162)
	FORMATAstc6x5UnormBlock                    = Format(163)
	FORMATAstc6x5SrgbBlock                     = Format(164)
	FORMATAstc6x6UnormBlock                    = Format(165)
	FORMATAstc6x6SrgbBlock                     = Format(166)
	FORMATAstc8x5UnormBlock                    = Format(167)
	FORMATAstc8x5SrgbBlock                     = Format(168)
	FORMATAstc8x6UnormBlock                    = Format(169)
	FORMATAstc8x6SrgbBlock                     = Format(170)
	FORMATAstc8x8UnormBlock                    = Format(171)
	FORMATAstc8x8SrgbBlock                     = Format(172)
	FORMATAstc10x5UnormBlock                   = Format(173)
	FORMATAstc10x5SrgbBlock                    = Format(174)
	FORMATAstc10x6UnormBlock                   = Format(175)
	FORMATAstc10x6SrgbBlock                    = Format(176)
	FORMATAstc10x8UnormBlock                   = Format(177)
	FORMATAstc10x8SrgbBlock                    = Format(178)
	FORMATAstc10x10UnormBlock                  = Format(179)
	FORMATAstc10x10SrgbBlock                   = Format(180)
	FORMATAstc12x10UnormBlock                  = Format(181)
	FORMATAstc12x10SrgbBlock                   = Format(182)
	FORMATAstc12x12UnormBlock                  = Format(183)
	FORMATAstc12x12SrgbBlock                   = Format(184)
	FORMATG8b8g8r8422Unorm                     = Format(1000156000)
	FORMATB8g8r8g8422Unorm                     = Format(1000156001)
	FORMATG8B8R83plane420Unorm                 = Format(1000156002)
	FORMATG8B8r82plane420Unorm                 = Format(1000156003)
	FORMATG8B8R83plane422Unorm                 = Format(1000156004)
	FORMATG8B8r82plane422Unorm                 = Format(1000156005)
	FORMATG8B8R83plane444Unorm                 = Format(1000156006)
	FORMATR10x6UnormPack16                     = Format(1000156007)
	FORMATR10x6g10x6Unorm2pack16               = Format(1000156008)
	FORMATR10x6g10x6b10x6a10x6Unorm4pack16     = Format(1000156009)
	FORMATG10x6b10x6g10x6r10x6422Unorm4pack16  = Format(1000156010)
	FORMATB10x6g10x6r10x6g10x6422Unorm4pack16  = Format(1000156011)
	FORMATG10x6B10x6R10x63plane420Unorm3pack16 = Format(1000156012)
	FORMATG10x6B10x6r10x62plane420Unorm3pack16 = Format(1000156013)
	FORMATG10x6B10x6R10x63plane422Unorm3pack16 = Format(1000156014)
	FORMATG10x6B10x6r10x62plane422Unorm3pack16 = Format(1000156015)
	FORMATG10x6B10x6R10x63plane444Unorm3pack16 = Format(1000156016)
	FORMATR12x4UnormPack16                     = Format(1000156017)
	FORMATR12x4g12x4Unorm2pack16               = Format(1000156018)
	FORMATR12x4g12x4b12x4a12x4Unorm4pack16     = Format(1000156019)
	FORMATG12x4b12x4g12x4r12x4422Unorm4pack16  = Format(1000156020)
	FORMATB12x4g12x4r12x4g12x4422Unorm4pack16  = Format(1000156021)
	FORMATG12x4B12x4R12x43plane420Unorm3pack16 = Format(1000156022)
	FORMATG12x4B12x4r12x42plane420Unorm3pack16 = Format(1000156023)
	FORMATG12x4B12x4R12x43plane422Unorm3pack16 = Format(1000156024)
	FORMATG12x4B12x4r12x42plane422Unorm3pack16 = Format(1000156025)
	FORMATG12x4B12x4R12x43plane444Unorm3pack16 = Format(1000156026)
	FORMATG16b16g16r16422Unorm                 = Format(1000156027)
	FORMATB16g16r16g16422Unorm                 = Format(1000156028)
	FORMATG16B16R163plane420Unorm              = Format(1000156029)
	FORMATG16B16r162plane420Unorm              = Format(1000156030)
	FORMATG16B16R163plane422Unorm              = Format(1000156031)
	FORMATG16B16r162plane422Unorm              = Format(1000156032)
	FORMATG16B16R163plane444Unorm              = Format(1000156033)
	FORMATPvrtc12bppUnormBlockImg              = Format(1000054000)
	FORMATPvrtc14bppUnormBlockImg              = Format(1000054001)
	FORMATPvrtc22bppUnormBlockImg              = Format(1000054002)
	FORMATPvrtc24bppUnormBlockImg              = Format(1000054003)
	FORMATPvrtc12bppSrgbBlockImg               = Format(1000054004)
	FORMATPvrtc14bppSrgbBlockImg               = Format(1000054005)
	FORMATPvrtc22bppSrgbBlockImg               = Format(1000054006)
	FORMATPvrtc24bppSrgbBlockImg               = Format(1000054007)
	FORMATAstc4x4SfloatBlockExt                = Format(1000066000)
	FORMATAstc5x4SfloatBlockExt                = Format(1000066001)
	FORMATAstc5x5SfloatBlockExt                = Format(1000066002)
	FORMATAstc6x5SfloatBlockExt                = Format(1000066003)
	FORMATAstc6x6SfloatBlockExt                = Format(1000066004)
	FORMATAstc8x5SfloatBlockExt                = Format(1000066005)
	FORMATAstc8x6SfloatBlockExt                = Format(1000066006)
	FORMATAstc8x8SfloatBlockExt                = Format(1000066007)
	FORMATAstc10x5SfloatBlockExt               = Format(1000066008)
	FORMATAstc10x6SfloatBlockExt               = Format(1000066009)
	FORMATAstc10x8SfloatBlockExt               = Format(1000066010)
	FORMATAstc10x10SfloatBlockExt              = Format(1000066011)
	FORMATAstc12x10SfloatBlockExt              = Format(1000066012)
	FORMATAstc12x12SfloatBlockExt              = Format(1000066013)
)

type PhysicalDeviceType int32

const (
	PHYSICALDeviceTypeOther         = PhysicalDeviceType(0)
	PHYSICALDeviceTypeIntegratedGpu = PhysicalDeviceType(1)
	PHYSICALDeviceTypeDiscreteGpu   = PhysicalDeviceType(2)
	PHYSICALDeviceTypeVirtualGpu    = PhysicalDeviceType(3)
	PHYSICALDeviceTypeCpu           = PhysicalDeviceType(4)
)

type ImageLayout int32

const (
	IMAGELayoutUndefined                             = ImageLayout(0)
	IMAGELayoutGeneral                               = ImageLayout(1)
	IMAGELayoutColorAttachmentOptimal                = ImageLayout(2)
	IMAGELayoutDepthStencilAttachmentOptimal         = ImageLayout(3)
	IMAGELayoutDepthStencilReadOnlyOptimal           = ImageLayout(4)
	IMAGELayoutShaderReadOnlyOptimal                 = ImageLayout(5)
	IMAGELayoutTransferSrcOptimal                    = ImageLayout(6)
	IMAGELayoutTransferDstOptimal                    = ImageLayout(7)
	IMAGELayoutPreinitialized                        = ImageLayout(8)
	IMAGELayoutDepthReadOnlyStencilAttachmentOptimal = ImageLayout(1000117000)
	IMAGELayoutDepthAttachmentStencilReadOnlyOptimal = ImageLayout(1000117001)
	IMAGELayoutPresentSrcKhr                         = ImageLayout(1000001002)
	IMAGELayoutSharedPresentKhr                      = ImageLayout(1000111000)
	IMAGELayoutShadingRateOptimalNv                  = ImageLayout(1000164003)
	IMAGELayoutFragmentDensityMapOptimalExt          = ImageLayout(1000218000)
)

type VertexInputRate int32

const (
	VERTEXInputRateVertex   = VertexInputRate(0)
	VERTEXInputRateInstance = VertexInputRate(1)
)

type SamplerAddressMode int32

const (
	SAMPLERAddressModeRepeat            = SamplerAddressMode(0)
	SAMPLERAddressModeMirroredRepeat    = SamplerAddressMode(1)
	SAMPLERAddressModeClampToEdge       = SamplerAddressMode(2)
	SAMPLERAddressModeClampToBorder     = SamplerAddressMode(3)
	SAMPLERAddressModeMirrorClampToEdge = SamplerAddressMode(4)
)

type DescriptorType int32

const (
	DESCRIPTORTypeSampler                 = DescriptorType(0)
	DESCRIPTORTypeCombinedImageSampler    = DescriptorType(1)
	DESCRIPTORTypeSampledImage            = DescriptorType(2)
	DESCRIPTORTypeStorageImage            = DescriptorType(3)
	DESCRIPTORTypeUniformTexelBuffer      = DescriptorType(4)
	DESCRIPTORTypeStorageTexelBuffer      = DescriptorType(5)
	DESCRIPTORTypeUniformBuffer           = DescriptorType(6)
	DESCRIPTORTypeStorageBuffer           = DescriptorType(7)
	DESCRIPTORTypeUniformBufferDynamic    = DescriptorType(8)
	DESCRIPTORTypeStorageBufferDynamic    = DescriptorType(9)
	DESCRIPTORTypeInputAttachment         = DescriptorType(10)
	DESCRIPTORTypeInlineUniformBlockExt   = DescriptorType(1000138000)
	DESCRIPTORTypeAccelerationStructureNv = DescriptorType(1000165000)
)

type ImageUsageFlags int32

const (
	IMAGEUsageTransferSrcBit            = ImageUsageFlags(0x1)
	IMAGEUsageTransferDstBit            = ImageUsageFlags(0x2)
	IMAGEUsageSampledBit                = ImageUsageFlags(0x4)
	IMAGEUsageStorageBit                = ImageUsageFlags(0x8)
	IMAGEUsageColorAttachmentBit        = ImageUsageFlags(0x10)
	IMAGEUsageDepthStencilAttachmentBit = ImageUsageFlags(0x20)
	IMAGEUsageTransientAttachmentBit    = ImageUsageFlags(0x40)
	IMAGEUsageInputAttachmentBit        = ImageUsageFlags(0x80)
	IMAGEUsageShadingRateImageBitNv     = ImageUsageFlags(0x100)
	IMAGEUsageFragmentDensityMapBitExt  = ImageUsageFlags(0x200)
)

type QueueFlags int32

const (
	QUEUEGraphicsBit      = QueueFlags(0x1)
	QUEUEComputeBit       = QueueFlags(0x2)
	QUEUETransferBit      = QueueFlags(0x4)
	QUEUESparseBindingBit = QueueFlags(0x8)
	QUEUEProtectedBit     = QueueFlags(0x10)
)

type PipelineStageFlags int32

const (
	PIPELINEStageTopOfPipeBit                    = PipelineStageFlags(0x1)
	PIPELINEStageDrawIndirectBit                 = PipelineStageFlags(0x2)
	PIPELINEStageVertexInputBit                  = PipelineStageFlags(0x4)
	PIPELINEStageVertexShaderBit                 = PipelineStageFlags(0x8)
	PIPELINEStageTessellationControlShaderBit    = PipelineStageFlags(0x10)
	PIPELINEStageTessellationEvaluationShaderBit = PipelineStageFlags(0x20)
	PIPELINEStageGeometryShaderBit               = PipelineStageFlags(0x40)
	PIPELINEStageFragmentShaderBit               = PipelineStageFlags(0x80)
	PIPELINEStageEarlyFragmentTestsBit           = PipelineStageFlags(0x100)
	PIPELINEStageLateFragmentTestsBit            = PipelineStageFlags(0x200)
	PIPELINEStageColorAttachmentOutputBit        = PipelineStageFlags(0x400)
	PIPELINEStageComputeShaderBit                = PipelineStageFlags(0x800)
	PIPELINEStageTransferBit                     = PipelineStageFlags(0x1000)
	PIPELINEStageBottomOfPipeBit                 = PipelineStageFlags(0x2000)
	PIPELINEStageHostBit                         = PipelineStageFlags(0x4000)
	PIPELINEStageAllGraphicsBit                  = PipelineStageFlags(0x8000)
	PIPELINEStageAllCommandsBit                  = PipelineStageFlags(0x10000)
	PIPELINEStageTransformFeedbackBitExt         = PipelineStageFlags(0x1000000)
	PIPELINEStageConditionalRenderingBitExt      = PipelineStageFlags(0x40000)
	PIPELINEStageCommandProcessBitNvx            = PipelineStageFlags(0x20000)
	PIPELINEStageShadingRateImageBitNv           = PipelineStageFlags(0x400000)
	PIPELINEStageRayTracingShaderBitNv           = PipelineStageFlags(0x200000)
	PIPELINEStageAccelerationStructureBuildBitNv = PipelineStageFlags(0x2000000)
	PIPELINEStageTaskShaderBitNv                 = PipelineStageFlags(0x80000)
	PIPELINEStageMeshShaderBitNv                 = PipelineStageFlags(0x100000)
	PIPELINEStageFragmentDensityProcessBitExt    = PipelineStageFlags(0x800000)
)

type ImageAspectFlags int32

const (
	IMAGEAspectColorBit           = ImageAspectFlags(0x1)
	IMAGEAspectDepthBit           = ImageAspectFlags(0x2)
	IMAGEAspectStencilBit         = ImageAspectFlags(0x4)
	IMAGEAspectMetadataBit        = ImageAspectFlags(0x8)
	IMAGEAspectPlane0Bit          = ImageAspectFlags(0x10)
	IMAGEAspectPlane1Bit          = ImageAspectFlags(0x20)
	IMAGEAspectPlane2Bit          = ImageAspectFlags(0x40)
	IMAGEAspectMemoryPlane0BitExt = ImageAspectFlags(0x80)
	IMAGEAspectMemoryPlane1BitExt = ImageAspectFlags(0x100)
	IMAGEAspectMemoryPlane2BitExt = ImageAspectFlags(0x200)
	IMAGEAspectMemoryPlane3BitExt = ImageAspectFlags(0x400)
)

type BufferUsageFlags int32

const (
	BUFFERUsageTransferSrcBit                       = BufferUsageFlags(0x1)
	BUFFERUsageTransferDstBit                       = BufferUsageFlags(0x2)
	BUFFERUsageUniformTexelBufferBit                = BufferUsageFlags(0x4)
	BUFFERUsageStorageTexelBufferBit                = BufferUsageFlags(0x8)
	BUFFERUsageUniformBufferBit                     = BufferUsageFlags(0x10)
	BUFFERUsageStorageBufferBit                     = BufferUsageFlags(0x20)
	BUFFERUsageIndexBufferBit                       = BufferUsageFlags(0x40)
	BUFFERUsageVertexBufferBit                      = BufferUsageFlags(0x80)
	BUFFERUsageIndirectBufferBit                    = BufferUsageFlags(0x100)
	BUFFERUsageTransformFeedbackBufferBitExt        = BufferUsageFlags(0x800)
	BUFFERUsageTransformFeedbackCounterBufferBitExt = BufferUsageFlags(0x1000)
	BUFFERUsageConditionalRenderingBitExt           = BufferUsageFlags(0x200)
	BUFFERUsageRayTracingBitNv                      = BufferUsageFlags(0x400)
	BUFFERUsageShaderDeviceAddressBitExt            = BufferUsageFlags(0x20000)
)

type ShaderStageFlags int32

const (
	SHADERStageVertexBit                 = ShaderStageFlags(0x1)
	SHADERStageTessellationControlBit    = ShaderStageFlags(0x2)
	SHADERStageTessellationEvaluationBit = ShaderStageFlags(0x4)
	SHADERStageGeometryBit               = ShaderStageFlags(0x8)
	SHADERStageFragmentBit               = ShaderStageFlags(0x10)
	SHADERStageComputeBit                = ShaderStageFlags(0x20)
	SHADERStageAllGraphics               = ShaderStageFlags(0x1f)
	SHADERStageAll                       = ShaderStageFlags(0x7fffffff)
	SHADERStageRaygenBitNv               = ShaderStageFlags(0x100)
	SHADERStageAnyHitBitNv               = ShaderStageFlags(0x200)
	SHADERStageClosestHitBitNv           = ShaderStageFlags(0x400)
	SHADERStageMissBitNv                 = ShaderStageFlags(0x800)
	SHADERStageIntersectionBitNv         = ShaderStageFlags(0x1000)
	SHADERStageCallableBitNv             = ShaderStageFlags(0x2000)
	SHADERStageTaskBitNv                 = ShaderStageFlags(0x40)
	SHADERStageMeshBitNv                 = ShaderStageFlags(0x80)
)

type DescriptorBindingFlagBitsEXT int32

const (
	DESCRIPTORBindingUpdateAfterBindBitExt          = DescriptorBindingFlagBitsEXT(1)
	DESCRIPTORBindingUpdateUnusedWhilePendingBitExt = DescriptorBindingFlagBitsEXT(2)
	DESCRIPTORBindingPartiallyBoundBitExt           = DescriptorBindingFlagBitsEXT(4)
	DESCRIPTORBindingVariableDescriptorCountBitExt  = DescriptorBindingFlagBitsEXT(8)
)